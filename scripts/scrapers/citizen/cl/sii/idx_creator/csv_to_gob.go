package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	siiTruncatedRutList := make([]string, 30_000_000)
	f, err := os.Open("scripts/assets/citizen/sii/sii_person_index.gob")
	if err != nil {
		log.Fatalf("Error opening ruts.gob: %v", err)
	}
	err = gob.NewDecoder(f).Decode(&siiTruncatedRutList)
	if err != nil {
		log.Fatalf("Error decoding ruts.gob: %v", err)
	}

	servelRuts := make([]string, 30_000_000)
	f, err = os.Open("scripts/assets/citizen/servel/servel_person_index.gob")
	if err != nil {
		log.Fatalf("Error opening ruts.gob: %v", err)
	}
	err = gob.NewDecoder(f).Decode(&servelRuts)
	if err != nil {
		log.Fatalf("Error decoding ruts.gob: %v", err)
	}
	var maxRut int
	if len(siiTruncatedRutList) > len(servelRuts) {
		maxRut = len(siiTruncatedRutList)
	} else {
		maxRut = len(servelRuts)
	}

	fmt.Printf("RUT máximo: %d\n", maxRut)
	fmt.Printf("Creando lista de personas...\n")
	// Crear un mapa para almacenar los datos procesados
	personList := make([]string, maxRut+1)
	for run := 0; run < len(servelRuts); run++ {
		name := servelRuts[run]
		personList[run] = name
	}

	for run := 0; run < len(siiTruncatedRutList); run++ {
		name := siiTruncatedRutList[run]
		personList[run] = name
	}

	// write to a single file
	indexFile, err := os.Create("scripts/assets/citizen/cl_partial_person_index.gob")
	if err != nil {
		panic(err)
	}
	defer indexFile.Close()
	encoder := gob.NewEncoder(indexFile)
	err = encoder.Encode(personList)
	if err != nil {
		panic(err)
	}

	voidSlotsCount := 0
	fillSlotsCount := 0
	for _, name := range personList {
		if name == "" {
			voidSlotsCount++
		} else {
			fillSlotsCount++
		}
	}

	fmt.Printf("Void slots: %d\n", voidSlotsCount)
	fmt.Printf("Fill slots: %d\n", fillSlotsCount)
	fmt.Printf("Total slots: %d\n", len(personList))
}
