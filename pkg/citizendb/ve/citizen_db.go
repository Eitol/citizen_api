package ve

import (
	"context"
	_ "embed"
	"errors"
	"github.com/Eitol/citizen_api/pkg/citizendb/shared"
	"runtime"
	"time"

	"github.com/Eitol/citizen_api/pkg/citizendb/names"
	"github.com/Eitol/citizen_api/pkg/hash"
	"log"
)

//go:embed location_index.json
var locationMapStr []byte

func NewCitizenDB(dbFilePath, dbLocationsFilePath, dbNameFilePath string, idVsNameMap []string) (*DB, error) {
	if idVsNameMap == nil {
		log.Fatalf("idVSNamePath is nil")
	}
	startTime := time.Now()
	locIndex, err := loadLocationIndex()
	if err != nil {
		return nil, err
	}
	runtime.GC()
	log.Printf("Location index loaded in %v\n", time.Since(startTime))
	rdb := &DB{
		LocationIndex: locIndex,
		IDVsNameMap:   idVsNameMap,
	}
	startTime = time.Now()
	rdb.CitizenDB = make([][11]byte, 0)
	errLoadCitizenDB := loadCitizenDB(dbFilePath, &rdb.CitizenDB)
	log.Printf("Citizen DB loaded in %v\n", time.Since(startTime))
	runtime.GC()
	startTime = time.Now()
	rdb.CitizenLocationDB = make([]uint16, 0)
	errLoadCitizenLocationDB := loadCitizenLocationDB(dbLocationsFilePath, &rdb.CitizenLocationDB)
	log.Printf("Citizen location DB loaded in %v\n", time.Since(startTime))
	runtime.GC()
	var errLoadNameDB error
	if dbNameFilePath != "" {
		startTime = time.Now()
		rdb.CitizenNamesDB = make(map[uint32]uint32)
		errLoadNameDB = loadNameDB(dbNameFilePath, &rdb.CitizenNamesDB)
		runtime.GC()
		log.Printf("Citizen names DB loaded in %v\n", time.Since(startTime))
	}
	if errLoadCitizenDB != nil {
		return nil, errLoadCitizenDB
	}
	if errLoadNameDB != nil {
		return nil, errLoadNameDB
	}
	if errLoadCitizenLocationDB != nil {
		return nil, errLoadCitizenLocationDB
	}
	return rdb, nil
}

type DB struct {
	LocationIndex     map[ParishID]Location
	CitizenDB         [][11]byte
	CitizenLocationDB []uint16
	CitizenNamesDB    map[uint32]uint32
	IDVsNameMap       []string
}

func (v *DB) decodeName(encodedName [11]byte) string {
	return names.DecodeNamesFrom11Bytes(encodedName, v.IDVsNameMap)
}

func (v *DB) FindCitizenNameByDocumentIDFast(docID int) string {
	return v.decodeName(v.CitizenDB[docID])
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
	if citizen == [11]byte{} {
		return nil, shared.ErrNotFound
	}
	location := v.LocationIndex[ParishID(v.CitizenLocationDB[docID])]
	return &Citizen{
		FullName:   v.decodeName(citizen),
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
	id, ok := v.CitizenNamesDB[nameHash]
	if !ok {
		return nil, shared.ErrNotFound
	}
	citizen := v.CitizenDB[id]
	if citizen == [11]byte{} {
		return nil, shared.ErrNotFound
	}
	var citizens []Citizen
	location := v.LocationIndex[ParishID(v.CitizenLocationDB[id])]
	citizens = append(citizens, Citizen{
		FullName:   v.decodeName(citizen),
		Location:   location,
		DocumentID: int(id),
	})
	return citizens, nil
}
