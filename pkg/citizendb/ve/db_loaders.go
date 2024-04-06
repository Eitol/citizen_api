package ve

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

func loadLocationIndex() (map[ParishID]Location, error) {
	idx := make(map[ParishID]Location, 1142)
	err := json.Unmarshal(locationMapStr, &idx)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling location index: %w", err)
	}
	return idx, nil
}

func loadCitizenDB(dbFilePath string) ([]OptimizedCitizen, error) {
	fileReader, err := os.Open(dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading citizen DB file: %w", err)
	}
	dbr := make([]OptimizedCitizen, 30_000_001)
	gobDev := gob.NewDecoder(fileReader)
	err = gobDev.Decode(&dbr)
	if err != nil {
		return nil, fmt.Errorf("error decoding citizen DB file: %w", err)
	}
	return dbr, nil
}

func loadNameDB(dbNameFilePath string) (map[uint32][]uint32, error) {
	citizenNameDBBytes, err := os.Open(dbNameFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading citizen names DB file: %w", err)
	}
	citizenNameDB := make(map[uint32][]uint32, 20889561)
	err = gob.NewDecoder(citizenNameDBBytes).Decode(&citizenNameDB)
	if err != nil {
		return nil, fmt.Errorf("error decoding citizen names DB file: %w", err)
	}
	return citizenNameDB, nil
}
