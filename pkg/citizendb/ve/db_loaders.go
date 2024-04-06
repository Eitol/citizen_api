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

func loadCitizenDB(dbFilePath string, dbr *[][11]byte) error {
	fileReader, err := os.Open(dbFilePath)
	if err != nil {
		return fmt.Errorf("error reading citizen DB file: %w", err)
	}
	gobDev := gob.NewDecoder(fileReader)
	err = gobDev.Decode(dbr)
	if err != nil {
		return fmt.Errorf("error decoding citizen DB file: %w", err)
	}
	return nil
}

func loadCitizenLocationDB(dbFilePath string, dbr *[]uint16) error {
	fileReader, err := os.Open(dbFilePath)
	if err != nil {
		return fmt.Errorf("error reading citizen DB file: %w", err)
	}
	gobDev := gob.NewDecoder(fileReader)
	err = gobDev.Decode(dbr)
	if err != nil {
		return fmt.Errorf("error decoding citizen DB file: %w", err)
	}
	return nil
}

func loadNameDB(dbNameFilePath string, citizenNameDB *map[uint32]uint32) error {
	citizenNameDBBytes, err := os.Open(dbNameFilePath)
	if err != nil {
		return fmt.Errorf("error reading citizen names DB file: %w", err)
	}
	err = gob.NewDecoder(citizenNameDBBytes).Decode(&citizenNameDB)
	if err != nil {
		return fmt.Errorf("error decoding citizen names DB file: %w", err)
	}
	return nil
}
