package ve

import (
	"github.com/Eitol/citizen_api/pkg/citizendb/names"
	"testing"
)

func Test_db_FindCitizenByDocumentID(t *testing.T) {
	idVsNameDB, err := names.LoadIDVsNameDB("../../../scripts/assets/citizen/names_list.gob")
	if err != nil {
		t.Errorf("LoadIDVsNameDB() error = %v", err)
		return
	}
	v, err := NewCitizenDB(
		"../../../scripts/assets/citizen/ve_optimized_citizen.gob",
		"../../../scripts/assets/citizen/ve_optimized_names.gob",
		idVsNameDB,
	)
	if err != nil {
		t.Errorf("NewCitizenDB() error = %v", err)
		return
	}

	got, err := v.FindCitizenByDocumentID(nil, 20300961)
	if err != nil {
		t.Errorf("FindCitizenByDocumentID() error = %v", err)
		return
	}
	if got.FullName != "HECTOR LUIS OLIVEROS LEON" {
		t.Errorf("FindCitizenByDocumentID() got = %v, want %v", got.FullName, "HECTOR LUIS OLIVEROS LEON")
	}

	gots, err := v.FindCitizenByName(nil, "HECTOR LUIS OLIVEROS LEON")
	if err != nil {
		t.Errorf("FindCitizenByName() error = %v", err)
		return
	}
	if gots[0].DocumentID != 20300961 {
		t.Errorf("FindCitizenByName() got = %v, want %v", gots[0].FullName, "HECTOR LUIS OLIVEROS LEON")
	}
}
