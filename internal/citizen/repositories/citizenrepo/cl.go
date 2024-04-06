package citizenrepo

import (
	"context"
	"errors"

	"github.com/Eitol/citizen_api/internal/citizen/domain"
	rututils "github.com/Eitol/rut"
)

func (m *MultiCountryCitizenRepository) findByCLRut(ctx context.Context, docID string) ([]domain.FindCitizenResult, error) {
	var citizens []domain.FindCitizenResult
	clCitizen, errCl := m.findByCLRutInRepo(ctx, docID)
	if errCl == nil {
		veCitizen, err := m.veDB.FindCitizenByName(ctx, clCitizen.Citizen.Name)
		if err == nil {
			for _, v := range veCitizen {
				citizens = append(citizens, domain.FindCitizenResult{
					Citizen:   m.adaptVeCitizen(v),
					MatchType: domain.MatchTypeByName,
				})
			}
		}
		citizens = append(citizens, *clCitizen)
	}
	return citizens, nil
}

func (m *MultiCountryCitizenRepository) findByCLRutInRepo(ctx context.Context, rut string) (*domain.FindCitizenResult, error) {
	parsedRut, err := rututils.Parse(rut)
	if err != nil {
		return nil, errors.Join(ErrInvalidCLRut, err)
	}
	clName, err := m.clDB.FindCitizenNameByRun(ctx, parsedRut.Run)
	if err != nil {
		return nil, err
	}
	return &domain.FindCitizenResult{
		Citizen: &domain.Citizen{
			Name: clName,
			Documents: []domain.DocumentID{
				{
					Number: parsedRut.Rut,
					Location: domain.Location{
						Country: "CL",
					},
				},
			},
		},
		MatchType: domain.MatchTypeByDocumentID,
	}, nil
}
