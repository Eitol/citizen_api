package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
	"github.com/Eitol/citizen_api/pkg/hash"
	"github.com/Eitol/citizen_api/pkg/strutils"
	"github.com/Eitol/cne_scraper/cne"
	"os"
	"path/filepath"
)

type veDBUnifier struct {
	locationMapPath string
	seniatFilesDir  string
	cneDBFile       string
	unificatedDB    string
	namesDB         string
}

func (u *veDBUnifier) loadLocationMap() map[ve.State]map[ve.Municipality]map[ve.Parish]ve.ParishLocation {
	b, err := os.ReadFile(u.locationMapPath)
	if err != nil {
		panic(err)
	}
	m := make(map[ve.State]map[ve.Municipality]map[ve.Parish]ve.ParishLocation)
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func (u *veDBUnifier) saveFile(s string, list interface{}) {
	fileDBName, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	defer func(fileDBName *os.File) {
		err := fileDBName.Close()
		if err != nil {
			panic(err)
		}
	}(fileDBName)
	encoder := gob.NewEncoder(fileDBName)
	err = encoder.Encode(list)
	if err != nil {
		panic(err)
	}
}

type Person struct {
	Cedula int
	Name   string
}

func (u *veDBUnifier) loadSeniatList() ([]Person, error) {
	gobDir := filepath.Join(u.seniatFilesDir)
	gobFiles, err := os.ReadDir(gobDir)
	if err != nil {
		return nil, err
	}
	personList := make([]Person, 0, 30_000_000)
	for _, gobFile := range gobFiles {
		filePath := filepath.Join(gobDir, gobFile.Name())
		b, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		var personListPart []Person
		err = gob.NewDecoder(bytes.NewReader(b)).Decode(&personListPart)
		if err != nil {
			return nil, err
		}
		personList = append(personList, personListPart...)
	}
	return personList, nil
}

func (u *veDBUnifier) loadCNEList() ([]cne.Person, error) {
	var personList []cne.Person
	file, err := os.ReadFile(u.cneDBFile)
	if err != nil {
		panic(err)
	}
	gobDecoderReader := bytes.NewReader(file)
	d := gob.NewDecoder(gobDecoderReader)
	err = d.Decode(&personList)
	if err != nil {
		panic(err)
	}
	return personList, nil
}

func (u *veDBUnifier) Start() {
	namesFixesCSVFile, err := os.ReadFile("scripts/assets/citizen/ve_citizen_weird_names.csv")
	if err != nil {
		panic(err)
	}
	csvFile, err := csv.NewReader(bytes.NewReader(namesFixesCSVFile)).ReadAll()
	if err != nil {
		panic(err)
	}
	nameFixesMap := make(map[string]string)
	for _, row := range csvFile {
		nameFixesMap[row[0]] = row[1]
	}
	locationMap := u.loadLocationMap()
	unificatedList := make([]ve.IndexedCitizen, 30_000_000)
	namesMap := make(map[uint32][]uint32)
	cneList, err := u.loadCNEList()
	if err != nil {
		panic(err)
	}
	seniatList, err := u.loadSeniatList()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loading CNE list: %d\n", len(cneList))
	for i := range cneList {
		name := strutils.RemoveAccents(cneList[i].Name)
		if cneList[i].Name == "" {
			continue
		}
		if newName, ok := nameFixesMap[name]; ok {
			name = newName
		}
		cneList[i].Name = name
		unificatedList[cneList[i].Cedula] = ve.IndexedCitizen{
			FullName:   cneList[i].Name,
			LocationID: ve.ParishID(locationMap[ve.State(cneList[i].State)][ve.Municipality(cneList[i].Municipality)][ve.Parish(cneList[i].Parish)].ID),
		}
	}
	fmt.Printf("Loading SENIAT list: %d\n", len(seniatList))
	for i := range seniatList {
		if unificatedList[seniatList[i].Cedula].FullName != "" {
			continue
		}
		unificatedList[seniatList[i].Cedula] = ve.IndexedCitizen{
			FullName: seniatList[i].Name,
		}
	}
	u.saveFile(u.unificatedDB, &unificatedList)
	fmt.Printf("Unificated list saved to %s\n", u.unificatedDB)
	fmt.Printf("Creating names map\n")
	for i := range unificatedList {
		if unificatedList[i].FullName == "" {
			continue
		}

		h := hash.HashFnv32(unificatedList[i].FullName)
		namesMap[h] = append(namesMap[h], uint32(i))
	}
	fmt.Printf("Saving names map\n")
	u.saveFile(u.namesDB, &namesMap)
}

func main() {
	u := veDBUnifier{
		locationMapPath: "scripts/assets/citizen/location_map.json",
		seniatFilesDir:  "scripts/assets/citizen/seniat",
		cneDBFile:       "scripts/assets/citizen/cne/cne.gob",
		unificatedDB:    "scripts/assets/citizen/ve_citizen.gob",
		namesDB:         "scripts/assets/citizen/ve_citizen_names.gob",
	}
	u.Start()
}
