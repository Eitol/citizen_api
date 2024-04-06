package domain

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type Location struct {
	Country      string
	State        string
	Municipality string
	Parish       string
	LocationID   string
	Coordinate   Coordinate
}

type DocumentID struct {
	Number   string
	Location Location
}

type Citizen struct {
	Name      string
	Documents []DocumentID
}

type MatchType int

const (
	MatchTypeByDocumentID MatchType = 1
	MatchTypeByName       MatchType = 2
)

type FindCitizenResult struct {
	Citizen   *Citizen
	MatchType MatchType
}
