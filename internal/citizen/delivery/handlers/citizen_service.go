package handlers

import (
	connect_go "connectrpc.com/connect"
	"context"
	"github.com/Eitol/citizen_api/internal/citizen/domain"
	v1 "github.com/Eitol/citizen_api/internal/gen/go/citizen/api/v1"
	"github.com/Eitol/citizen_api/internal/gen/go/citizen/api/v1/apiv1connect"
)

func NewCitizenHandler(
	findPersonByDocumentIDUseCase domain.FindPersonByDocumentIDUseCase,
) apiv1connect.CitizenServiceHandler {
	return &citizenHandler{
		findPersonByDocumentIDUseCase: findPersonByDocumentIDUseCase,
	}
}

type citizenHandler struct {
	findPersonByDocumentIDUseCase domain.FindPersonByDocumentIDUseCase
}

func (c *citizenHandler) FindCitizenByDocId(ctx context.Context, req *connect_go.Request[v1.FindCitizenByDocIdRequest]) (*connect_go.Response[v1.FindCitizenByDocIdResponse], error) {
	result, err := c.findPersonByDocumentIDUseCase.Execute(ctx, req.Msg.DocumentId)
	if err != nil {
		return nil, err
	}
	r := adaptFindCitizenResultToResponse(result)
	return r, nil
}
