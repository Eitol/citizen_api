package ve

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

func loadLocationIndex() (map[ParishID]Location, error) {
	idx := map[ParishID]Location{}
	err := json.Unmarshal(locationMapStr, &idx)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling location index: %w", err)
	}
	return idx, nil
}

func loadCitizenDB(dbFilePath string) ([]OptimizedCitizen, error) {
	fileReader, err := os.ReadFile(dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading citizen DB file: %w", err)
	}
	var dbr []OptimizedCitizen
	gobDev := gob.NewDecoder(bytes.NewReader(fileReader))
	err = gobDev.Decode(&dbr)
	if err != nil {
		return nil, fmt.Errorf("error decoding citizen DB file: %w", err)
	}
	return dbr, nil
}

func loadNameDB(dbNameFilePath string) (map[uint32][]uint32, error) {
	citizenNameDBBytes, err := os.ReadFile(dbNameFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading citizen names DB file: %w", err)
	}
	var citizenNameDB map[uint32][]uint32
	citizenNameDBBytesReader := bytes.NewReader(citizenNameDBBytes)
	err = gob.NewDecoder(citizenNameDBBytesReader).Decode(&citizenNameDB)
	if err != nil {
		return nil, fmt.Errorf("error decoding citizen names DB file: %w", err)
	}
	return citizenNameDB, nil
}

func loadIDVsNameMap(idVSNamePath string) (map[uint32]string, error) {
	idVSNameMap := map[uint32]string{}
	idVSNameBytes, err := os.ReadFile(idVSNamePath)
	if err != nil {
		return nil, fmt.Errorf("error reading ID vs Name map file: %w", err)
	}
	err = gob.NewDecoder(bytes.NewReader(idVSNameBytes)).Decode(&idVSNameMap)
	if err != nil {
		return nil, fmt.Errorf("error decoding ID vs Name map file: %w", err)
	}
	return idVSNameMap, nil
}
