package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	servelPersonList := make([]string, 0, 30_000_000)
	var siiPersonMap map[int]string
	fmt.Printf("Opening files...\n")
	servelPersonListFile, err := os.OpenFile("scripts/assets/citizen/servel/servel_person_index.gob", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decoding servel_person_index.gob...\n")
	err = gob.NewDecoder(servelPersonListFile).Decode(&servelPersonList)
	if err != nil {
		panic(err)
	}
	siiPersonListFile, err := os.OpenFile("scripts/assets/citizen/sii/cl_sii_map_db.gob", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decoding sii_person_index.gob...\n")
	err = gob.NewDecoder(siiPersonListFile).Decode(&siiPersonMap)
	if err != nil {
		panic(err)
	}

	servelPersonListFile.Close()
	siiPersonListFile.Close()
	maxRut := 0
	unifiedPersonMap := make(map[int]string, len(servelPersonList)+len(siiPersonMap))
	fmt.Printf("Unifying lists...\n")
	for i, person := range servelPersonList {
		if person == "" {
			continue
		}
		maxRut = i
		unifiedPersonMap[i] = person
	}

	fmt.Printf("Max RUT: %d\n", maxRut)
	fmt.Printf("Unifying sii map...\n")
	for i, person := range siiPersonMap {
		if person == "" {
			continue
		}
		if i > maxRut {
			maxRut = i
		}
		unifiedPersonMap[i] = person
	}

	unifiedPersonListFile, err := os.Create("scripts/assets/citizen/cl_unified_person_index.gob")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saving unified_person_index.gob...\n")
	err = gob.NewEncoder(unifiedPersonListFile).Encode(unifiedPersonMap)
	if err != nil {
		panic(err)
	}
	unifiedPersonListFile.Close()

	unifiedPersonList := make([]string, maxRut+1)
	for i, person := range unifiedPersonMap {
		unifiedPersonList[i] = person
	}

	unifiedPersonListFile, err = os.Create("scripts/assets/citizen/cl_unified_person_list.gob")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saving unified_person_list.gob...\n")
	err = gob.NewEncoder(unifiedPersonListFile).Encode(unifiedPersonList)
	if err != nil {
		panic(err)
	}
	unifiedPersonListFile.Close()
}
