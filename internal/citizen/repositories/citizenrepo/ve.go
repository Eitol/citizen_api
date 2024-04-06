package citizenrepo

import (
	"context"
	"errors"
	"strconv"

	"github.com/Eitol/citizen_api/internal/citizen/domain"
	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
)

func (m *MultiCountryCitizenRepository) findByVEDocumentID(ctx context.Context, docID string) (*domain.FindCitizenResult, error) {
	parsedDocID, err := strconv.Atoi(docID)
	if err != nil {
		return nil, errors.Join(ErrInvalidVEDocID, err)
	}
	veCitizen, err := m.veDB.FindCitizenByDocumentID(ctx, parsedDocID)
	if err != nil {
		return nil, err
	}

	c := m.adaptVeCitizen(*veCitizen)
	return &domain.FindCitizenResult{
		Citizen:   c,
		MatchType: domain.MatchTypeByDocumentID,
	}, nil
}

func (m *MultiCountryCitizenRepository) adaptVeCitizen(veCitizen ve.Citizen) *domain.Citizen {
	return &domain.Citizen{
		Name: veCitizen.FullName,
		Documents: []domain.DocumentID{
			{
				Number: strconv.Itoa(veCitizen.DocumentID),
				Location: domain.Location{
					Country:      "VE",
					State:        string(veCitizen.Location.State),
					Municipality: string(veCitizen.Location.Municipality),
					Parish:       string(veCitizen.Location.Parish),
					LocationID:   strconv.Itoa(int(veCitizen.Location.ParishID)),
					Coordinate: domain.Coordinate{
						Latitude:  veCitizen.Location.Coordinate.Latitude,
						Longitude: veCitizen.Location.Coordinate.Longitude,
					},
				},
			},
		},
	}
}
