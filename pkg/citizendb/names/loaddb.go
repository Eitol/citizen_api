package names

import (
	"encoding/gob"
	"os"
	"runtime"
)

func LoadIDVsNameDB(path string) ([]string, error) {
	idVsNameMap := make([]string, 2_900_000)
	f, err := os.Open(path)
	// gob decode idVsNameBytes into idVsNameMap
	err = gob.NewDecoder(f).Decode(&idVsNameMap)
	if err != nil {
		return nil, err
	}
	runtime.GC()
	return idVsNameMap, nil
}
