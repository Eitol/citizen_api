package ve

import (
	"context"
	_ "embed"
	"errors"
	"github.com/Eitol/citizen_api/pkg/citizendb/shared"

	"log"
	"sync"

	"github.com/Eitol/citizen_api/pkg/citizendb/names"
	"github.com/Eitol/citizen_api/pkg/hash"
)

//go:embed location_index.json
var locationMapStr []byte

func NewCitizenDB(dbFilePath, dbNameFilePath string, idVsNameMap []string) (*DB, error) {
	if idVsNameMap == nil {
		log.Fatalf("idVSNamePath is nil")
	}
	locIndex, err := loadLocationIndex()
	if err != nil {
		return nil, err
	}
	rdb := &DB{
		LocationIndex: locIndex,
		IDVsNameMap:   idVsNameMap,
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	var errLoadCitizenDB error
	go func() {
		defer wg.Done()
		rdb.CitizenDB, errLoadCitizenDB = loadCitizenDB(dbFilePath)
	}()
	var errLoadNameDB error
	if dbNameFilePath != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rdb.CitizenNamesDB, errLoadNameDB = loadNameDB(dbNameFilePath)
		}()
	}
	wg.Wait()
	if errLoadCitizenDB != nil {
		return nil, errLoadCitizenDB
	}
	if errLoadNameDB != nil {
		return nil, errLoadNameDB
	}
	return rdb, nil
}

type DB struct {
	LocationIndex  map[ParishID]Location
	CitizenDB      []OptimizedCitizen
	CitizenNamesDB map[uint32][]uint32
	IDVsNameMap    []string
}

func (v *DB) decodeName(encodedName [11]byte) string {
	return names.DecodeNamesFrom11Bytes(encodedName, v.IDVsNameMap)
}

func (v *DB) FindCitizenNameByDocumentIDFast(docID int) string {
	return v.decodeName(v.CitizenDB[docID].FullName)
}

func (v *DB) FindCitizenByDocumentID(ctx context.Context, docID int) (*Citizen, error) {
	if ctx.Err() != nil {
		return nil, errors.New("context error")
	}
	if docID < 0 {
		return nil, shared.ErrInvalidDocumentID
	}
	if docID >= len(v.CitizenDB) {
		return nil, shared.ErrOutOfRange
	}
	citizen := v.CitizenDB[docID]
	if citizen.FullName == [11]byte{} {
		return nil, shared.ErrNotFound
	}
	location := v.LocationIndex[ParishID(citizen.LocationID)]
	return &Citizen{
		FullName:   v.decodeName(citizen.FullName),
		Location:   location,
		DocumentID: docID,
	}, nil
}

func (v *DB) FindCitizenByName(ctx context.Context, name string) ([]Citizen, error) {
	if ctx.Err() != nil {
		return nil, errors.New("context error")
	}
	if v.CitizenNamesDB == nil {
		return nil, errors.New("name DB not provided")
	}
	nameHash := hash.HashFnv32(name)
	ids, ok := v.CitizenNamesDB[nameHash]
	if !ok {
		return nil, shared.ErrNotFound
	}
	citizens := make([]Citizen, 0, len(ids))
	for _, id := range ids {
		citizen := v.CitizenDB[id]
		if citizen.FullName == [11]byte{} {
			continue
		}
		location := v.LocationIndex[ParishID(citizen.LocationID)]
		citizens = append(citizens, Citizen{
			FullName:   v.decodeName(citizen.FullName),
			Location:   location,
			DocumentID: int(id),
		})
	}
	return citizens, nil
}
