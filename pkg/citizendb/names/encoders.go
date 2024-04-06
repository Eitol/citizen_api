package names

import "strings"

func EncodeNamesIn11Bytes(s string, m map[string]uint32) [11]byte {
	names := splitNameIn4Slices(s)
	ids := [4]uint32{0, 0, 0, 0}
	for i, name := range names {
		ids[i] = uint32(m[name])
	}
	return encode4IntsIn11Bytes(ids)
}

func splitNameIn4Slices(s string) [4]string {
	spName := strings.Split(s, " ")
	sp := make([]string, 0, len(spName))
	for _, s := range spName {
		if len(s) == 0 || s == "" || s == " " {
			continue
		}
		sp = append(sp, s)
	}
	if len(sp) == 0 {
		return [4]string{}
	}
	if len(sp) == 1 {
		return [4]string{sp[0]}
	}
	if len(sp) == 2 {
		return [4]string{sp[0], sp[1]}
	}
	if len(sp) == 3 {
		return [4]string{sp[0], sp[1], sp[2]}
	}
	if len(sp) == 4 {
		return [4]string{sp[0], sp[1], sp[2], sp[3]}
	}
	return [4]string{
		sp[0],
		strings.Join(sp[1:len(sp)-2], " "),
		sp[len(sp)-2],
		sp[len(sp)-1],
	}
}

func encode4IntsIn11Bytes(ids [4]uint32) [11]byte {
	var result [11]byte

	// Codifica cada entero en los 22 bits y los asigna a la matriz de bytes de resultado
	for i, id := range ids {
		// Calcula el desplazamiento de bit para el entero actual
		shift := uint(22 * i)

		// Inserta los bits en la posición correcta del arreglo de bytes
		for bit := 0; bit < 22; bit++ {
			byteIndex := (shift + uint(bit)) / 8
			bitIndex := (shift + uint(bit)) % 8

			if id&(1<<bit) != 0 {
				result[byteIndex] |= 1 << bitIndex
			}
		}
	}

	return result
}
