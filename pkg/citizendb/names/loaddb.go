package names

import (
	"bytes"
	"encoding/gob"
	"os"
)

func LoadIDVsNameDB(path string) ([]string, error) {
	idVsNameMap := make([]string, 30_000_000)
	f, err := os.ReadFile(path)
	// gob decode idVsNameBytes into idVsNameMap
	err = gob.NewDecoder(bytes.NewReader(f)).Decode(&idVsNameMap)
	if err != nil {
		return nil, err
	}
	return idVsNameMap, nil
}
