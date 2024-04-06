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
	outBuf := make([]byte, 0, 1024*1024+1)
	for i := 0; true; i++ {
		var c *ve.Citizen
		c, err = v.FindCitizenByDocumentID(nil, i)
		if err != nil {
			continue
		}
		name := strutils.RemoveAccents(c.FullName)
		outBuf = append(outBuf, name...)
	}
	err = os.WriteFile("scripts/assets/citizen/ve_citizen_names.txt", outBuf, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
