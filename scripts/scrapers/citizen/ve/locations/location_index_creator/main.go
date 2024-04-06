package main

import (
	"encoding/json"
	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
	"os"
)

func main() {
	b, err := os.ReadFile("scripts/assets/citizen/location_map.json")
	if err != nil {
		panic(err)
	}
	var l map[ve.State]map[ve.Municipality]map[ve.Parish]ve.ParishLocation
	err = json.Unmarshal(b, &l)
	if err != nil {
		panic(err)
	}
	locationIndex := map[ve.ParishID]ve.Location{}
	for state, municipalities := range l {
		for municipality, parishes := range municipalities {
			for parish, id := range parishes {
				locationIndex[ve.ParishID(id.ID)] = ve.Location{
					State:        state,
					Municipality: municipality,
					Parish:       parish,
					Locality:     ve.Locality(id.Locality),
					ParishID:     ve.ParishID(id.ID),
					Coordinate: ve.Coordinate{
						Latitude:  id.Lat,
						Longitude: id.Lon,
					},
				}
			}
		}
	}

	b, err = json.MarshalIndent(locationIndex, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("scripts/assets/citizen/location_index.json", b, 0644)
}
