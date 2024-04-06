package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/Eitol/gosii"
	"github.com/Eitol/gosii/pkg"
)

const (
	defaultStartIDX   = 1
	endID             = 30000000
	numThreads        = 100
	chunkSize         = 1000
	successOutputDir  = "scripts/assets/citizen/sii/rut_success"
	failedOutputDir   = "scripts/assets/citizen/sii/rut_failed"
	latestIDXFileName = "scripts/assets/citizen/sii/rut_latest.idx"
)

type Stats struct {
	TotalCount        int
	LatestAvgTime     float64
	LatestAvgAttempts float64
	StartTime         time.Time
}

type siiScraper struct {
	personMutex       sync.Mutex
	failedPersonList  []int
	failedPersonMutex sync.Mutex
	stats             Stats
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatalf("Error closing file: %v", err)
	}
}

func readLatestIDXFile() int {
	var latestID int
	file, err := os.Open(latestIDXFileName)
	if err != nil {
		return defaultStartIDX
	}
	defer closeFile(file)
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&latestID)
	if err != nil {
		return defaultStartIDX
	}
	return latestID
}

func saveLatestIDXFile(latestID int) {
	file, err := os.Create(latestIDXFileName)
	if err != nil {
		return
	}
	defer closeFile(file)
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(latestID)
	if err != nil {
		return
	}
}

func OnNewCaptcha(captcha *gosii.Captcha) {
	fmt.Println("New captcha!:" + captcha.Solution)
}

func (s *siiScraper) ScrapRuts() {
	s.stats.StartTime = time.Now()
	runList := make([]string, 30_000_000)
	f, err := os.Open("scripts/assets/citizen/cl_partial_person_index.gob")
	if err != nil {
		log.Fatalf("Error opening ruts.gob: %v", err)
	}
	err = gob.NewDecoder(f).Decode(&runList)
	if err != nil {
		log.Fatalf("Error decoding ruts.gob: %v", err)
	}

	var personList []gosii.Citizen
	startIDX := readLatestIDXFile()
	log.Printf("Starting from ID: %d\n", startIDX)
	var wg sync.WaitGroup
	idChan := make(chan int)
	err = os.MkdirAll(successOutputDir, 0777)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}
	err = os.MkdirAll(failedOutputDir, 0777)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	// Iniciar las goroutines
	ssiClient := gosii.NewClient(&gosii.Opts{OnNewCaptcha: OnNewCaptcha})
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.listenRUNChannel(idChan, ssiClient, &personList)
		}()
		startDelay := 1 * time.Second
		time.Sleep(startDelay)
	}
	latestChunkProcessedID := 0
	beginTime := time.Now()
	lenRunDB := len(runList)
	// Enviar los IDs a las goroutines
	for run := startIDX + 1; run <= endID; run++ {
		if run < lenRunDB && runList[run] != "" {
			continue
		}
		idChan <- run
		if run%1000 == 0 {
			fmt.Printf("ID: %d | success %d | failed | %d \n", run, len(personList), len(s.failedPersonList))
			s.printStats()
		}
		if len(personList) >= chunkSize {
			totalTime := time.Since(beginTime)
			fmt.Printf("Saving chunk | ID: %d | time: %v | total %d \n", run, totalTime, run-latestChunkProcessedID)
			s.personMutex.Lock()
			s.failedPersonMutex.Lock()

			savePersonList(personList, s.failedPersonList, run)
			saveLatestIDXFile(run)
			// clear lists
			personList = personList[:0]
			s.failedPersonList = s.failedPersonList[:0]

			s.personMutex.Unlock()
			s.failedPersonMutex.Unlock()
			latestChunkProcessedID = run
			beginTime = time.Now()
		}
	}

	// Cerrar el canal y esperar a que todas las goroutines terminen
	close(idChan)
	wg.Wait()
}

func (s *siiScraper) listenRUNChannel(idChan chan int, ssiClient gosii.Client, personList *[]gosii.Citizen) {
	for id := range idChan {
		s.processRUN(id, ssiClient, personList)
	}
}

func (s *siiScraper) updateStats(meta *gosii.RequestMetadata) {
	s.stats.TotalCount = meta.TotalCount
	s.stats.LatestAvgTime = float64(meta.AvgTime)
	s.stats.LatestAvgAttempts = float64(meta.Attempts)
}

func (s *siiScraper) printStats() {
	fmt.Printf("#Reqs: %d | AvgTime(Secs): %d\n", s.stats.TotalCount, time.Duration(s.stats.LatestAvgTime)/time.Second)
	fmt.Printf("Elapsed time: %v\n", time.Since(s.stats.StartTime))
}

func (s *siiScraper) processRUN(run int, ssiClient gosii.Client, personList *[]gosii.Citizen) {
	rut, person, meta, isNotFound, err := findCitizenByRUN(run, ssiClient)
	if meta != nil {
		s.updateStats(meta)
	}
	if isNotFound {
		return
	}

	if err != nil || person == nil || person.Name == "" {
		s.failedPersonMutex.Lock()
		s.failedPersonList = append(s.failedPersonList, run)
		s.failedPersonMutex.Unlock()
		return
	}
	if person.Rut == "" {
		person.Rut = rut
	}
	if person.Name != "" {
		s.personMutex.Lock()
		*personList = append(*personList, *person)
		s.personMutex.Unlock()
	}
}

func findCitizenByRUN(run int, ssiClient gosii.Client) (string, *gosii.Citizen, *gosii.RequestMetadata, bool, error) {
	dv := pkg.GetRutDv(run)
	rut := fmt.Sprintf("%d-%s", run, dv)
	citizen, meta, err := ssiClient.GetNameByRUT(rut)
	if err != nil {
		if errors.Is(err, gosii.ErrNotFound) || err.Error() == "not found" {
			return rut, nil, meta, true, nil
		}
		return rut, nil, meta, false, err
	}
	return rut, citizen, meta, false, nil
}

func savePersonList(personList []gosii.Citizen, failedPersonList []int, id int) {
	err := saveList(personList, id, successOutputDir)
	if err != nil {
		log.Fatalf("Error saving person list: %v", err)
	}
	err = saveList(failedPersonList, id, failedOutputDir)
	if err != nil {
		log.Fatalf("Error saving failed person list: %v", err)
	}
}

func saveList[V any](personList V, id int, path string) error {
	fileName := filepath.Join(path, strconv.Itoa(id)+".gob")
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer closeFile(file)
	err = gob.NewEncoder(file).Encode(personList)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	s := siiScraper{}
	s.ScrapRuts()
}
