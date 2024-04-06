package ve

type State string
type Municipality string
type Parish string
type ParishID int32
type Locality string

type IndexedCitizen struct {
	FullName   string
	LocationID ParishID
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type Location struct {
	State        State        `json:"state"`
	Municipality Municipality `json:"municipality"`
	Parish       Parish       `json:"parish"`
	Locality     Locality     `json:"locality"`
	ParishID     ParishID     `json:"id"`
	Coordinate   Coordinate   `json:"coordinate"`
}

type ParishLocation struct {
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	ID       int     `json:"id"`
	Locality string  `json:"locality"`
}

type Citizen struct {
	FullName   string
	DocumentID int
	Location   Location
}

type OptimizedCitizen struct {
	FullName   [11]byte
	LocationID uint16
}
