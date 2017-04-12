package main

import (
	"math"
	"unicode"
)

func runeWidth(r rune) (width int, ashex bool) {
	if unicode.In(r, unicode.Me, unicode.Mn, unicode.Cf) ||
		(r >= 0x1160 && r <= 0x11FF) ||
		!unicode.IsPrint(r) {

		return int(math.Log2(float64(r))/4) + 3, true
	} else if unicode.In(r, &wideRanges) {
		return 2, false
	}
	return 1, false
}
