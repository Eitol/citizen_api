package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/Eitol/citizen_api/pkg/strutils"
	"os"

	"github.com/Eitol/citizen_api/pkg/citizendb/names"
	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
	"github.com/Eitol/citizen_api/pkg/hash"
)

func main() {
	namesMap := map[string]uint32{}
	namesMapFile, err := os.ReadFile("scripts/assets/citizen/names_map.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewDecoder(bytes.NewReader(namesMapFile)).Decode(&namesMap)
	if err != nil {
		panic(err)
	}
	v := make([]ve.IndexedCitizen, 30_000_000)
	vf, err := os.ReadFile("scripts/assets/citizen/ve_citizen.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewDecoder(bytes.NewReader(vf)).Decode(&v)
	optimizedDB := make([]ve.OptimizedCitizen, 30_000_000)
	optimizedNamesDB := make([][11]byte, 30_000_000)
	optimizedLocationsDB := make([]uint16, 30_000_000)
	maxRut := 0
	for i := 1; i < 30_000_000; i++ {
		if i%1_000_000 == 0 {
			fmt.Println("Processed", i, "names")
		}
		citizen := v[i]
		if len(citizen.FullName) == 0 {
			continue
		}
		if i > maxRut {
			maxRut = i
		}
		cleanFullName := strutils.RemoveAccents(citizen.FullName)
		v[i].FullName = cleanFullName
		encodedName := names.EncodeNamesIn11Bytes(cleanFullName, namesMap)
		optimizedDB[i] = ve.OptimizedCitizen{
			FullName:   encodedName,
			LocationID: uint16(citizen.LocationID),
		}
		optimizedNamesDB[i] = encodedName
		optimizedLocationsDB[i] = uint16(citizen.LocationID)
	}
	fmt.Println("Max RUT", maxRut)
	namesDB := make(map[uint32][]uint32, maxRut+1)
	namesUniqueDB := make(map[uint32]uint32, maxRut+1)
	for i := 0; i < maxRut; i++ {
		c := v[i]
		hashOfTheName := hash.HashFnv32(c.FullName)
		_, ok := namesDB[hashOfTheName]
		if !ok {
			namesDB[hashOfTheName] = []uint32{uint32(i)}
			namesUniqueDB[hashOfTheName] = uint32(i)
		} else {
			namesDB[hashOfTheName] = append(namesDB[hashOfTheName], uint32(i))
			namesUniqueDB[hashOfTheName] = uint32(i)
		}
	}
	optimizedDBFile, err := os.Create("scripts/assets/citizen/ve_optimized_citizen.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(optimizedDBFile).Encode(optimizedDB)
	if err != nil {
		panic(err)
	}

	optimizedLocationsDBFile, err := os.Create("scripts/assets/citizen/ve_optimized_locations.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(optimizedLocationsDBFile).Encode(optimizedLocationsDB)
	if err != nil {
		panic(err)
	}

	optimizedNamesDBFile, err := os.Create("scripts/assets/citizen/ve_optimized_names_list.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(optimizedNamesDBFile).Encode(optimizedNamesDB)
	if err != nil {
		panic(err)
	}

	namesDBFile, err := os.Create("scripts/assets/citizen/ve_optimized_names.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(namesDBFile).Encode(namesDB)
	if err != nil {
		panic(err)
	}

	namesUniqueDBFile, err := os.Create("scripts/assets/citizen/ve_optimized_names_unique.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(namesUniqueDBFile).Encode(namesUniqueDB)
	if err != nil {
		panic(err)
	}
}
