package cl

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/Eitol/citizen_api/pkg/citizendb/shared"
	"os"
	"runtime"
)

type DB struct {
	index []string
}

func NewDB(dbPath string) (*DB, error) {
	index := make([]string, 30_000_000)
	// Load index from dbPath
	indexFile, err := os.Open(dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening index file: %w", err)
	}
	err = gob.NewDecoder(indexFile).Decode(&index)
	if err != nil {
		return nil, fmt.Errorf("error decoding index file: %w", err)
	}
	runtime.GC()
	err = indexFile.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing index file: %w", err)
	}
	return &DB{
		index: index,
	}, nil
}

func (v *DB) FindCitizenNameByRun(ctx context.Context, run int) (string, error) {
	if ctx.Err() != nil {
		return "", fmt.Errorf("context error: %w", ctx.Err())
	}
	if run < 0 || run >= len(v.index) {
		return "", shared.ErrOutOfRange
	}
	name := v.index[run]
	if name == "" {
		return "", shared.ErrNotFound
	}
	return v.index[run], nil
}
