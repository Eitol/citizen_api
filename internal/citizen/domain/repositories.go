package domain

import "context"

type CitizenRepository interface {
	FindByDocumentID(ctx context.Context, docID string) ([]FindCitizenResult, error)
}
