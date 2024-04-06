package main

import (
	"encoding/gob"
	"fmt"
	"github.com/Eitol/gosii"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	dir := "scripts/assets/citizen/sii/rut_success"
	// list all gob files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var out []gosii.Citizen
	outMap := map[int]string{}
	for _, filePath := range files {
		if filePath.IsDir() {
			continue
		}
		// read the gob filePath
		var file *os.File
		file, err = os.Open(filepath.Join(dir, filePath.Name()))
		if err != nil {
			panic(err)
		}
		var gobFile []gosii.Citizen
		err = gob.NewDecoder(file).Decode(&gobFile)
		if err != nil {
			panic(err)
		}
		out = append(out, gobFile...)
		err = file.Close()
		if err != nil {
			panic(err)
		}
		for _, c := range gobFile {
			if len(c.Rut) < 2 {
				continue
			}
			run, err := strconv.Atoi(c.Rut[:len(c.Rut)-2])
			if err != nil {
				panic(err)
			}
			outMap[run] = c.Name
		}
	}
	fmt.Printf("Cantidad de personas: %d\n", len(out))
	// save the out
	file, err := os.Create("scripts/assets/citizen/sii/cl_sii_list_db.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(file).Encode(out)
	if err != nil {
		panic(err)
	}
	err = file.Close()

	file, err = os.Create("scripts/assets/citizen/sii/cl_sii_map_db.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(file).Encode(outMap)
	if err != nil {
		panic(err)
	}
}
