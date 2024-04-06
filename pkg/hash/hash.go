package hash

import "hash/fnv"

func HashFnv32(str string) uint32 {
	hasher := fnv.New32()
	_, err := hasher.Write([]byte(str))
	if err != nil {
		return 0
	}
	return hasher.Sum32()
}
