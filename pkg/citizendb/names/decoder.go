package names

import "strings"

func DecodeNamesFrom11Bytes(b [11]byte, m []string) string {
	ids := decode4IntsFrom11Bytes(b)
	var names = make([]string, 0, 4)
	for _, id := range ids {
		if id == 0 {
			continue
		}
		names = append(names, m[uint32(id)-1])
	}
	return strings.Trim(strings.Join(names, " "), " ")
}

func decode4IntsFrom11Bytes(encodedName [11]byte) [4]uint32 {
	var result [4]uint32

	// Decodes each integer from the 22 bits and assign it to the result int array
	for i := 0; i < 4; i++ {
		shift := uint(22 * i)
		var id uint32
		for bit := 0; bit < 22; bit++ {
			byteIndex := (shift + uint(bit)) / 8
			bitIndex := (shift + uint(bit)) % 8

			if encodedName[byteIndex]&(1<<bitIndex) != 0 {
				id |= 1 << bit
			}
		}
		result[i] = id
	}
	return result
}
