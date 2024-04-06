package strutils

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
)

// RemoveAccents transforms the string, removing accents and preserving "Ñ"
func RemoveAccents(s string) string {
	s = strings.ToUpper(s)
	// Reemplace temporalmente "ñ" y "Ñ" por "ӈ" y "Ӊ" que son caracteres latinos sin usar en español
	s = strings.ReplaceAll(s, "Ñ", "Ӊ")

	transformer := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		runes.Remove(runes.In(unicode.Me)),
		norm.NFC,
	)
	accentRemoved, _, _ := transform.String(transformer, s)
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ӈӉ]+")
	clean := reg.ReplaceAllString(accentRemoved, "")

	// Revertir la "ñ" y "Ñ" a su forma original
	clean = strings.ReplaceAll(clean, "Ӊ", "Ñ")
	clean = strings.ReplaceAll(clean, "0", "O")
	clean = strings.Trim(clean, " ")
	// reemplace los espacios múltiples por uno solo
	sp := strings.Split(clean, " ")
	outSP := ""
	for _, sS := range sp {
		if sS != "" {
			outSP += sS + " "
		}
	}
	clean = strings.Trim(outSP, " ")
	return clean
}
