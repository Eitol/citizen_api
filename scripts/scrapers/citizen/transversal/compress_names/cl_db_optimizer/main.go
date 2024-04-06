package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/Eitol/citizen_api/pkg/strutils"
	"os"
	"sort"
	"strings"

	"github.com/Eitol/citizen_api/pkg/citizendb/ve"
)

func splitName(s string) []string {
	spName := strings.Split(s, " ")
	sp := make([]string, 0, len(spName))
	for _, name := range spName {
		if len(name) == 0 || strings.ReplaceAll(name, " ", "") == "" {
			continue
		}
		name = strutils.RemoveAccents(name)
		sp = append(sp, name)
	}
	if len(sp) == 0 {
		return nil
	}
	if len(sp) == 1 {
		return []string{sp[0]}
	}
	if len(sp) == 2 {
		return []string{sp[0], sp[1]}
	}
	if len(sp) == 3 {
		return []string{sp[0], sp[1], sp[2]}
	}
	if len(sp) == 4 {
		return []string{sp[0], sp[1], sp[2], sp[3]}
	}
	return []string{
		sp[0],
		strings.Join(sp[1:len(sp)-2], " "),
		sp[len(sp)-2],
		sp[len(sp)-1],
	}
}

func main() {
	v := make([]ve.IndexedCitizen, 30_000_000)
	vf, err := os.ReadFile("scripts/assets/citizen/ve_citizen.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewDecoder(bytes.NewReader(vf)).Decode(&v)

	clMap := map[int]string{}
	clFile, err := os.ReadFile("scripts/assets/citizen/cl_unified_person_index.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewDecoder(bytes.NewReader(clFile)).Decode(&clMap)
	if err != nil {
		panic(err)
	}
	nameSet := make(map[string]struct{}, 4_000_000)
	for i := 0; i < 30_000_000; i++ {
		if i%1_000_000 == 0 {
			fmt.Println("Processed", i, "names")
		}
		c := v[i]
		spName := splitName(c.FullName)
		for _, s := range spName {
			nameSet[s] = struct{}{}
		}
		clName, ok := clMap[i]
		if ok {
			spName = splitName(clName)
			for _, s := range spName {
				nameSet[s] = struct{}{}
			}
		}
	}
	nameList := make([]string, 0, len(nameSet))
	for k := range nameSet {
		nameList = append(nameList, k)
	}
	sort.Strings(nameList)

	namesMap := make(map[string]uint32, len(nameList))
	for i, name := range nameList {
		namesMap[name] = uint32(i + 1)
	}

	file, err := os.Create("scripts/assets/citizen/names_list.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(file).Encode(nameList)
	if err != nil {
		panic(err)
	}
	file.Close()
	file, err = os.Create("scripts/assets/citizen/names_map.gob")
	if err != nil {
		panic(err)
	}
	err = gob.NewEncoder(file).Encode(namesMap)
}
