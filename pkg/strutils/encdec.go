package strutils

import (
	"bytes"
	"fmt"
)

type EncDec struct {
	charMap    map[rune]int
	reverseMap map[int]rune
	escapeChar int
}

func NewEncDec() *EncDec {
	// Mapeo inicial de caracteres comunes, la Ñ se incluye por su uso en español.
	charMap := map[rune]int{
		'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5,
		'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10,
		'K': 11, 'L': 12, 'M': 13, 'N': 14, 'Ñ': 15,
		// Usamos valores negativos como marcadores para caracteres que requieren el carácter de escape.
		'O': -1, 'P': -2, 'Q': -3, 'R': -4, 'S': -5,
		'T': -6, 'U': -7, 'V': -8, 'W': -9, 'X': -10,
		'Y': -11, 'Z': -12,
	}

	reverseMap := make(map[int]rune)
	for k, v := range charMap {
		reverseMap[v] = k
	}

	return &EncDec{charMap, reverseMap, 0}
}

func (e *EncDec) Encode(s string) []byte {
	var buffer bytes.Buffer
	escapeNext := false
	for _, r := range s {
		value, exists := e.charMap[r]
		if !exists {
			continue // Ignore characters not in our map
		}

		if value < 0 { // Characters requiring the escape character
			if !escapeNext {
				buffer.WriteByte(byte(e.escapeChar << 4))
				escapeNext = true
			}
			buffer.WriteByte(byte(-value))
		} else {
			if escapeNext {
				escapeNext = false
			} else {
				buffer.WriteByte(byte(value << 4))
			}
		}
	}
	return buffer.Bytes()
}

func (e *EncDec) Decode(bytes []byte) string {
	var result []rune
	escapeNext := false
	for _, b := range bytes {
		if escapeNext {
			result = append(result, e.reverseMap[-int(b)])
			escapeNext = false
		} else if int(b>>4) == e.escapeChar {
			escapeNext = true
		} else {
			result = append(result, e.reverseMap[int(b>>4)])
		}
	}
	return string(result)
}

func main() {
	encDec := NewEncDec()
	encoded := encDec.Encode("HOLA MUNDOÑ")
	fmt.Println("Encoded:", encoded)
	decoded := encDec.Decode(encoded)
	fmt.Println("Decoded:", decoded)
}
