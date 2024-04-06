package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"

	"github.com/Eitol/cne_scraper/cne"
)

func main() {
	var personList []cne.Person
	file, err := os.ReadFile("scripts/assets/citizen/cne/cne.gob")
	if err != nil {
		panic(err)
	}
	gobDecoderReader := bytes.NewReader(file)
	d := gob.NewDecoder(gobDecoderReader)
	err = d.Decode(&personList)
	if err != nil {
		panic(err)
	}
	fmt.Print(personList[0])
}
