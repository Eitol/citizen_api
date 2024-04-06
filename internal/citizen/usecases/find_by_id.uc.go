package usecases

import (
	"context"
	"github.com/Eitol/citizen_api/internal/citizen/domain"
)

type findByIDUC struct {
	repo domain.CitizenRepository
}

func NewFindByIDUC(repo domain.CitizenRepository) domain.FindPersonByDocumentIDUseCase {
	return &findByIDUC{repo: repo}
}

func (f *findByIDUC) Execute(ctx context.Context, id string) ([]domain.FindCitizenResult, error) {
	return f.repo.FindByDocumentID(ctx, id)
}
