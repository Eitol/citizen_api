package handlers

import (
	connect_go "connectrpc.com/connect"

	"github.com/Eitol/citizen_api/internal/citizen/domain"
	v1 "github.com/Eitol/citizen_api/internal/gen/go/citizen/api/v1"
)

func adaptFindCitizenResultToResponse(result []domain.FindCitizenResult) *connect_go.Response[v1.FindCitizenByDocIdResponse] {
	results := make([]*v1.FindCitizenResult, len(result))
	for i, r := range result {
		results[i] = &v1.FindCitizenResult{
			Citizen: &v1.Citizen{
				Name:      r.Citizen.Name,
				Documents: adaptDocuments(r.Citizen.Documents),
			},
			MatchType: adaptMatchType(r.MatchType),
		}
	}
	return &connect_go.Response[v1.FindCitizenByDocIdResponse]{
		Msg: &v1.FindCitizenByDocIdResponse{
			Results: results,
		},
	}
}

func adaptMatchType(matchType domain.MatchType) v1.MatchType {
	switch matchType {
	case domain.MatchTypeByDocumentID:
		return v1.MatchType_MATCH_TYPE_BY_DOCUMENT_ID
	case domain.MatchTypeByName:
		return v1.MatchType_MATCH_TYPE_BY_NAME
	default:
		return v1.MatchType_MATCH_TYPE_UNSPECIFIED
	}
}

func adaptDocuments(documents []domain.DocumentID) []*v1.DocumentID {
	docs := make([]*v1.DocumentID, len(documents))
	for i, d := range documents {
		docs[i] = &v1.DocumentID{
			Number: d.Number,
			Location: &v1.Location{
				Country:      d.Location.Country,
				State:        d.Location.State,
				Municipality: d.Location.Municipality,
				Parish:       d.Location.Parish,
				LocationId:   d.Location.LocationID,
				Latitude:     d.Location.Coordinate.Latitude,
				Longitude:    d.Location.Coordinate.Longitude,
			},
		}
	}
	return docs
}
