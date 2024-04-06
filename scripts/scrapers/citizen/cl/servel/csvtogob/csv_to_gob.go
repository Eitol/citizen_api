package main

import (
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	baseDir := "scripts/assets/citizen/servel/"
	// Abrir el archivo CSV
	csvFile, err := os.Open(baseDir + "servel.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// Crear un lector de CSV
	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = ',' // Define el delimitador, en este caso la coma
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Número de registros: %d\n", len(records))

	rutMap := make(map[int]uint32, len(records)+1)
	maxRut := 0
	for i := 0; i < len(records); i++ {
		record := records[i]
		// Suponiendo que cada línea del CSV tiene el formato "Nombre,RUT"
		rut := record[1][0 : len(record[1])-2] // Eliminar el dígito verificador
		rut = strings.ReplaceAll(rut, ".", "") // Eliminar puntos del RUT

		// Convertir RUT a entero
		rutInt, err := strconv.Atoi(rut)
		if err != nil {
			panic(err)
		}
		rutMap[i] = uint32(rutInt)
		if rutInt > maxRut {
			maxRut = rutInt
		}
	}

	fmt.Printf("RUT máximo: %d\n", maxRut)
	fmt.Printf("Creando lista de personas...\n")
	// Crear un mapa para almacenar los datos procesados
	personList := make([]string, maxRut+1)

	for i := 0; i < len(records); i++ {
		// Suponiendo que cada línea del CSV tiene el formato "Nombre,RUT"
		name := records[i][0]
		rut := rutMap[i]
		// Almacenar en el mapa
		personList[rut] = name
	}

	fmt.Printf("Guardando lista de personas en formato GOB...\n")

	// Crear/abrir un archivo para escribir el mapa en formato GOB
	file, err := os.Create(baseDir + "servel_person_index.gob")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Crear un codificador GOB y codificar el mapa en el archivo
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(personList); err != nil {
		panic(err)
	}
}
