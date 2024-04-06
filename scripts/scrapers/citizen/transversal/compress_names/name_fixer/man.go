package main

import (
	"fmt"
	"github.com/Eitol/citizen_api/pkg/citizendb/names"
	"github.com/Eitol/citizen_api/pkg/strutils"
	"os"

	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
)

func main() {
	idVsNameDB, err := names.LoadIDVsNameDB("../../../scripts/assets/citizen/names_list.gob")
	if err != nil {
		fmt.Println("LoadIDVsNameDB() error = %v", err)
		return
	}
	v, err := ve.NewCitizenDB(
		"scripts/assets/citizen/ve_citizen.gob",
		"",
		idVsNameDB,
	)
	if err != nil {
		panic(err)
	}
	var weirdNames []string
	for i := 0; i < 29_000_000; i++ {
		if i%1_000_000 == 0 {
			fmt.Println("Processed", i, "names")
		}
		name := strutils.RemoveAccents(v.FindCitizenNameByDocumentIDFast(i))
		if name == "" {
			continue
		}
		for _, r := range name {
			if r == ' ' {
				r = 0
			} else if r == 'Ñ' {
				r = 1
			} else {
				r -= 63
				if r < 0 || r > 29 {
					weirdNames = append(weirdNames, name)
					break
				}
			}
		}
	}
	csvText := ""
	for _, name := range weirdNames {
		csvText += name + ",\n"
	}
	err = os.WriteFile("scripts/assets/citizen/ve_citizen_weird_names.csv", []byte(csvText), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
