package main

import (
	"math"
	"unicode"
)

func runeWidth(r rune) int {
	if unicode.In(r, unicode.Me, unicode.Mn, unicode.Cf) ||
		(r >= 0x1160 && r <= 0x11FF) ||
		!unicode.IsPrint(r) {

		return int(math.Log2(float64(r))/4) + 3 // <XXXX>
	} else if unicode.In(r, &wideRanges) {
		return 2
	}
	return 1
}
