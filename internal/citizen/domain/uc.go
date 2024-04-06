package domain

import "context"

type FindPersonByDocumentIDUseCase interface {
	Execute(ctx context.Context, id string) ([]FindCitizenResult, error)
}
