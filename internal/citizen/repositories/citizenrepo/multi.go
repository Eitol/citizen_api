package citizenrepo

import (
	"context"
	"errors"

	"github.com/Eitol/citizen_api/internal/citizen/domain"
	"github.com/Eitol/citizen_api/pkg/citizendb/cl"
	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
)

var (
	ErrInvalidCLRut   = errors.New("invalid chilean rut")
	ErrInvalidVEDocID = errors.New("invalid venezuelan document id")
	ErrNotFound       = errors.New("citizen not found")
)

type MultiCountryCitizenRepository struct {
	clDB *cl.DB
	veDB *ve.DB
}

func NewMultiCountryCitizenRepository(clDB *cl.DB, veDB *ve.DB) domain.CitizenRepository {
	return &MultiCountryCitizenRepository{
		clDB: clDB,
		veDB: veDB,
	}
}

func (m *MultiCountryCitizenRepository) FindByDocumentID(ctx context.Context, docID string) ([]domain.FindCitizenResult, error) {
	clCitizen, errCl := m.findByCLRut(ctx, docID)
	veCitizen, errVe := m.findByVEDocumentID(ctx, docID)
	if errCl != nil && errVe != nil {
		return nil, ErrNotFound
	}
	var citizens []domain.FindCitizenResult
	if errCl == nil && len(clCitizen) > 0 {
		citizens = append(citizens, clCitizen...)
	}
	if errVe == nil && veCitizen != nil {
		citizens = append(citizens, *veCitizen)
	}
	return citizens, nil
}
