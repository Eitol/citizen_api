package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/Eitol/citizen_api/pkg/citizendb/names"
	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
	"github.com/Eitol/citizen_api/pkg/strutils"
	"os"
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
	charHistogram := [30]int{}
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
		for _, r := range name {
			if r == ' ' {
				r = 0
			} else if r == 'Ñ' {
				r = 1
			} else {
				r -= 63
				if r < 0 || r > 29 {
					break
				}
			}
			charHistogram[r]++
		}
	}
	charHistogramMap := make(map[rune]int)
	for i, vv := range charHistogram {
		var r rune
		if i == 0 {
			r = '-'
		} else if i == 1 {
			r = 'Ñ'
		} else {
			r = rune(i) + 63
		}
		charHistogramMap[r] = vv
	}
	outJSON, err := json.Marshal(charHistogramMap)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("scripts/assets/citizen/ve_citizen_names_char_histogram.txt", outJSON, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
