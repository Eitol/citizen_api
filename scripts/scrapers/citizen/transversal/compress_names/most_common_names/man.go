package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/Eitol/citizen_api/pkg/citizendb/names"
	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
	"github.com/Eitol/citizen_api/pkg/strutils"
	"os"
	"sort"
)

func main() {
	namesFixesCSVFile, err := os.ReadFile("scripts/assets/citizen/ve_citizen_weird_names.csv")
	if err != nil {
		panic(err)
	}
	idVsNameDB, err := names.LoadIDVsNameDB("../../../scripts/assets/citizen/names_list.gob")
	if err != nil {
		fmt.Printf("LoadIDVsNameDB() error = %v\n", err)
		return
	}
	csvFile, err := csv.NewReader(bytes.NewReader(namesFixesCSVFile)).ReadAll()
	v, err := ve.NewCitizenDB(
		"scripts/assets/citizen/ve_citizen.gob",
		"",
		idVsNameDB,
	)
	if err != nil {
		panic(err)
	}
	nameFixesMap := make(map[string]string)
	for _, row := range csvFile {
		nameFixesMap[row[0]] = row[1]
	}
	nameHistogram := make(map[string]int, 30_000_000)
	for i := 0; i < 29_000_000; i++ {
		if i%1_000_000 == 0 {
			fmt.Println("Processed", i, "names")
		}
		name := strutils.RemoveAccents(v.FindCitizenNameByDocumentIDFast(i))
		if name == "" {
			continue
		}
		if newName, ok := nameFixesMap[name]; ok {
			name = newName
		}
		_, ok := nameHistogram[name]
		if !ok {
			nameHistogram[name] = 1
		} else {
			nameHistogram[name] = nameHistogram[name] + 1
		}
	}
	sortedNames := make([]string, 0, len(nameHistogram))
	sortedValues := make([]int, 0, len(nameHistogram))
	for i, val := range nameHistogram {
		sortedNames = append(sortedNames, i)
		sortedValues = append(sortedValues, val)
	}
	sort.Slice(sortedNames, func(i, j int) bool {
		return sortedValues[i] > sortedValues[j]
	})
	sort.Slice(sortedValues, func(i, j int) bool {
		return sortedValues[i] > sortedValues[j]
	})
	sortedCSVFile, err := os.Create("scripts/assets/citizen/ve_citizen_sorted_names.csv")
	if err != nil {
		panic(err)
	}
	sortedCSV := csv.NewWriter(sortedCSVFile)

	for i := 0; i < len(sortedNames); i++ {
		sortedCSV.Write([]string{sortedNames[i], fmt.Sprint(sortedValues[i])})
	}
	sortedCSV.Flush()
	sortedCSVFile.Close()
	fmt.Println("Done")
}
