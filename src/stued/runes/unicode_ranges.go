package runes

import (
    "unicode"
)

// something like...
//
// grep '^[^;]*;[FW]' EastAsianWidth.txt
// | sed 's/;[WF][ \t]*#.*$//'
// | sed 's/^\(....\)\.\.\(....\)/{Lo: \1, Hi: \2, Stride: 1}/'
// | sed 's/^\(....\)$/{Lo: \1, Hi: \1, Stride: 1}/'
// | sed 's/^\(.*\)\.\.\(.*\)/{Lo: \1, Hi: \2, Stride: 1}/'
// | sed 's/^\([^,]*\)$/{Lo: \1, Hi: \1, Stride: 1}/'
// | sed 's/: /: 0x/g'
//

var wideRanges = unicode.RangeTable{
	R16: []unicode.Range16{
		{Lo: 0x1100, Hi: 0x115F, Stride: 0x1},
		{Lo: 0x231A, Hi: 0x231B, Stride: 0x1},
		{Lo: 0x2329, Hi: 0x2329, Stride: 0x1},
		{Lo: 0x232A, Hi: 0x232A, Stride: 0x1},
		{Lo: 0x23E9, Hi: 0x23EC, Stride: 0x1},
		{Lo: 0x23F0, Hi: 0x23F0, Stride: 0x1},
		{Lo: 0x23F3, Hi: 0x23F3, Stride: 0x1},
		{Lo: 0x25FD, Hi: 0x25FE, Stride: 0x1},
		{Lo: 0x2614, Hi: 0x2615, Stride: 0x1},
		{Lo: 0x2648, Hi: 0x2653, Stride: 0x1},
		{Lo: 0x267F, Hi: 0x267F, Stride: 0x1},
		{Lo: 0x2693, Hi: 0x2693, Stride: 0x1},
		{Lo: 0x26A1, Hi: 0x26A1, Stride: 0x1},
		{Lo: 0x26AA, Hi: 0x26AB, Stride: 0x1},
		{Lo: 0x26BD, Hi: 0x26BE, Stride: 0x1},
		{Lo: 0x26C4, Hi: 0x26C5, Stride: 0x1},
		{Lo: 0x26CE, Hi: 0x26CE, Stride: 0x1},
		{Lo: 0x26D4, Hi: 0x26D4, Stride: 0x1},
		{Lo: 0x26EA, Hi: 0x26EA, Stride: 0x1},
		{Lo: 0x26F2, Hi: 0x26F3, Stride: 0x1},
		{Lo: 0x26F5, Hi: 0x26F5, Stride: 0x1},
		{Lo: 0x26FA, Hi: 0x26FA, Stride: 0x1},
		{Lo: 0x26FD, Hi: 0x26FD, Stride: 0x1},
		{Lo: 0x2705, Hi: 0x2705, Stride: 0x1},
		{Lo: 0x270A, Hi: 0x270B, Stride: 0x1},
		{Lo: 0x2728, Hi: 0x2728, Stride: 0x1},
		{Lo: 0x274C, Hi: 0x274C, Stride: 0x1},
		{Lo: 0x274E, Hi: 0x274E, Stride: 0x1},
		{Lo: 0x2753, Hi: 0x2755, Stride: 0x1},
		{Lo: 0x2757, Hi: 0x2757, Stride: 0x1},
		{Lo: 0x2795, Hi: 0x2797, Stride: 0x1},
		{Lo: 0x27B0, Hi: 0x27B0, Stride: 0x1},
		{Lo: 0x27BF, Hi: 0x27BF, Stride: 0x1},
		{Lo: 0x2B1B, Hi: 0x2B1C, Stride: 0x1},
		{Lo: 0x2B50, Hi: 0x2B50, Stride: 0x1},
		{Lo: 0x2B55, Hi: 0x2B55, Stride: 0x1},
		{Lo: 0x2E80, Hi: 0x2E99, Stride: 0x1},
		{Lo: 0x2E9B, Hi: 0x2EF3, Stride: 0x1},
		{Lo: 0x2F00, Hi: 0x2FD5, Stride: 0x1},
		{Lo: 0x2FF0, Hi: 0x2FFB, Stride: 0x1},
		{Lo: 0x3000, Hi: 0x3000, Stride: 0x1},
		{Lo: 0x3001, Hi: 0x3003, Stride: 0x1},
		{Lo: 0x3004, Hi: 0x3004, Stride: 0x1},
		{Lo: 0x3005, Hi: 0x3005, Stride: 0x1},
		{Lo: 0x3006, Hi: 0x3006, Stride: 0x1},
		{Lo: 0x3007, Hi: 0x3007, Stride: 0x1},
		{Lo: 0x3008, Hi: 0x3008, Stride: 0x1},
		{Lo: 0x3009, Hi: 0x3009, Stride: 0x1},
		{Lo: 0x300A, Hi: 0x300A, Stride: 0x1},
		{Lo: 0x300B, Hi: 0x300B, Stride: 0x1},
		{Lo: 0x300C, Hi: 0x300C, Stride: 0x1},
		{Lo: 0x300D, Hi: 0x300D, Stride: 0x1},
		{Lo: 0x300E, Hi: 0x300E, Stride: 0x1},
		{Lo: 0x300F, Hi: 0x300F, Stride: 0x1},
		{Lo: 0x3010, Hi: 0x3010, Stride: 0x1},
		{Lo: 0x3011, Hi: 0x3011, Stride: 0x1},
		{Lo: 0x3012, Hi: 0x3013, Stride: 0x1},
		{Lo: 0x3014, Hi: 0x3014, Stride: 0x1},
		{Lo: 0x3015, Hi: 0x3015, Stride: 0x1},
		{Lo: 0x3016, Hi: 0x3016, Stride: 0x1},
		{Lo: 0x3017, Hi: 0x3017, Stride: 0x1},
		{Lo: 0x3018, Hi: 0x3018, Stride: 0x1},
		{Lo: 0x3019, Hi: 0x3019, Stride: 0x1},
		{Lo: 0x301A, Hi: 0x301A, Stride: 0x1},
		{Lo: 0x301B, Hi: 0x301B, Stride: 0x1},
		{Lo: 0x301C, Hi: 0x301C, Stride: 0x1},
		{Lo: 0x301D, Hi: 0x301D, Stride: 0x1},
		{Lo: 0x301E, Hi: 0x301F, Stride: 0x1},
		{Lo: 0x3020, Hi: 0x3020, Stride: 0x1},
		{Lo: 0x3021, Hi: 0x3029, Stride: 0x1},
		{Lo: 0x302A, Hi: 0x302D, Stride: 0x1},
		{Lo: 0x302E, Hi: 0x302F, Stride: 0x1},
		{Lo: 0x3030, Hi: 0x3030, Stride: 0x1},
		{Lo: 0x3031, Hi: 0x3035, Stride: 0x1},
		{Lo: 0x3036, Hi: 0x3037, Stride: 0x1},
		{Lo: 0x3038, Hi: 0x303A, Stride: 0x1},
		{Lo: 0x303B, Hi: 0x303B, Stride: 0x1},
		{Lo: 0x303C, Hi: 0x303C, Stride: 0x1},
		{Lo: 0x303D, Hi: 0x303D, Stride: 0x1},
		{Lo: 0x303E, Hi: 0x303E, Stride: 0x1},
		{Lo: 0x3041, Hi: 0x3096, Stride: 0x1},
		{Lo: 0x3099, Hi: 0x309A, Stride: 0x1},
		{Lo: 0x309B, Hi: 0x309C, Stride: 0x1},
		{Lo: 0x309D, Hi: 0x309E, Stride: 0x1},
		{Lo: 0x309F, Hi: 0x309F, Stride: 0x1},
		{Lo: 0x30A0, Hi: 0x30A0, Stride: 0x1},
		{Lo: 0x30A1, Hi: 0x30FA, Stride: 0x1},
		{Lo: 0x30FB, Hi: 0x30FB, Stride: 0x1},
		{Lo: 0x30FC, Hi: 0x30FE, Stride: 0x1},
		{Lo: 0x30FF, Hi: 0x30FF, Stride: 0x1},
		{Lo: 0x3105, Hi: 0x312D, Stride: 0x1},
		{Lo: 0x3131, Hi: 0x318E, Stride: 0x1},
		{Lo: 0x3190, Hi: 0x3191, Stride: 0x1},
		{Lo: 0x3192, Hi: 0x3195, Stride: 0x1},
		{Lo: 0x3196, Hi: 0x319F, Stride: 0x1},
		{Lo: 0x31A0, Hi: 0x31BA, Stride: 0x1},
		{Lo: 0x31C0, Hi: 0x31E3, Stride: 0x1},
		{Lo: 0x31F0, Hi: 0x31FF, Stride: 0x1},
		{Lo: 0x3200, Hi: 0x321E, Stride: 0x1},
		{Lo: 0x3220, Hi: 0x3229, Stride: 0x1},
		{Lo: 0x322A, Hi: 0x3247, Stride: 0x1},
		{Lo: 0x3250, Hi: 0x3250, Stride: 0x1},
		{Lo: 0x3251, Hi: 0x325F, Stride: 0x1},
		{Lo: 0x3260, Hi: 0x327F, Stride: 0x1},
		{Lo: 0x3280, Hi: 0x3289, Stride: 0x1},
		{Lo: 0x328A, Hi: 0x32B0, Stride: 0x1},
		{Lo: 0x32B1, Hi: 0x32BF, Stride: 0x1},
		{Lo: 0x32C0, Hi: 0x32FE, Stride: 0x1},
		{Lo: 0x3300, Hi: 0x33FF, Stride: 0x1},
		{Lo: 0x3400, Hi: 0x4DB5, Stride: 0x1},
		{Lo: 0x4DB6, Hi: 0x4DBF, Stride: 0x1},
		{Lo: 0x4E00, Hi: 0x9FD5, Stride: 0x1},
		{Lo: 0x9FD6, Hi: 0x9FFF, Stride: 0x1},
		{Lo: 0xA000, Hi: 0xA014, Stride: 0x1},
		{Lo: 0xA015, Hi: 0xA015, Stride: 0x1},
		{Lo: 0xA016, Hi: 0xA48C, Stride: 0x1},
		{Lo: 0xA490, Hi: 0xA4C6, Stride: 0x1},
		{Lo: 0xA960, Hi: 0xA97C, Stride: 0x1},
		{Lo: 0xAC00, Hi: 0xD7A3, Stride: 0x1},
		{Lo: 0xF900, Hi: 0xFA6D, Stride: 0x1},
		{Lo: 0xFA6E, Hi: 0xFA6F, Stride: 0x1},
		{Lo: 0xFA70, Hi: 0xFAD9, Stride: 0x1},
		{Lo: 0xFADA, Hi: 0xFAFF, Stride: 0x1},
		{Lo: 0xFE10, Hi: 0xFE16, Stride: 0x1},
		{Lo: 0xFE17, Hi: 0xFE17, Stride: 0x1},
		{Lo: 0xFE18, Hi: 0xFE18, Stride: 0x1},
		{Lo: 0xFE19, Hi: 0xFE19, Stride: 0x1},
		{Lo: 0xFE30, Hi: 0xFE30, Stride: 0x1},
		{Lo: 0xFE31, Hi: 0xFE32, Stride: 0x1},
		{Lo: 0xFE33, Hi: 0xFE34, Stride: 0x1},
		{Lo: 0xFE35, Hi: 0xFE35, Stride: 0x1},
		{Lo: 0xFE36, Hi: 0xFE36, Stride: 0x1},
		{Lo: 0xFE37, Hi: 0xFE37, Stride: 0x1},
		{Lo: 0xFE38, Hi: 0xFE38, Stride: 0x1},
		{Lo: 0xFE39, Hi: 0xFE39, Stride: 0x1},
		{Lo: 0xFE3A, Hi: 0xFE3A, Stride: 0x1},
		{Lo: 0xFE3B, Hi: 0xFE3B, Stride: 0x1},
		{Lo: 0xFE3C, Hi: 0xFE3C, Stride: 0x1},
		{Lo: 0xFE3D, Hi: 0xFE3D, Stride: 0x1},
		{Lo: 0xFE3E, Hi: 0xFE3E, Stride: 0x1},
		{Lo: 0xFE3F, Hi: 0xFE3F, Stride: 0x1},
		{Lo: 0xFE40, Hi: 0xFE40, Stride: 0x1},
		{Lo: 0xFE41, Hi: 0xFE41, Stride: 0x1},
		{Lo: 0xFE42, Hi: 0xFE42, Stride: 0x1},
		{Lo: 0xFE43, Hi: 0xFE43, Stride: 0x1},
		{Lo: 0xFE44, Hi: 0xFE44, Stride: 0x1},
		{Lo: 0xFE45, Hi: 0xFE46, Stride: 0x1},
		{Lo: 0xFE47, Hi: 0xFE47, Stride: 0x1},
		{Lo: 0xFE48, Hi: 0xFE48, Stride: 0x1},
		{Lo: 0xFE49, Hi: 0xFE4C, Stride: 0x1},
		{Lo: 0xFE4D, Hi: 0xFE4F, Stride: 0x1},
		{Lo: 0xFE50, Hi: 0xFE52, Stride: 0x1},
		{Lo: 0xFE54, Hi: 0xFE57, Stride: 0x1},
		{Lo: 0xFE58, Hi: 0xFE58, Stride: 0x1},
		{Lo: 0xFE59, Hi: 0xFE59, Stride: 0x1},
		{Lo: 0xFE5A, Hi: 0xFE5A, Stride: 0x1},
		{Lo: 0xFE5B, Hi: 0xFE5B, Stride: 0x1},
		{Lo: 0xFE5C, Hi: 0xFE5C, Stride: 0x1},
		{Lo: 0xFE5D, Hi: 0xFE5D, Stride: 0x1},
		{Lo: 0xFE5E, Hi: 0xFE5E, Stride: 0x1},
		{Lo: 0xFE5F, Hi: 0xFE61, Stride: 0x1},
		{Lo: 0xFE62, Hi: 0xFE62, Stride: 0x1},
		{Lo: 0xFE63, Hi: 0xFE63, Stride: 0x1},
		{Lo: 0xFE64, Hi: 0xFE66, Stride: 0x1},
		{Lo: 0xFE68, Hi: 0xFE68, Stride: 0x1},
		{Lo: 0xFE69, Hi: 0xFE69, Stride: 0x1},
		{Lo: 0xFE6A, Hi: 0xFE6B, Stride: 0x1},
		{Lo: 0xFF01, Hi: 0xFF03, Stride: 0x1},
		{Lo: 0xFF04, Hi: 0xFF04, Stride: 0x1},
		{Lo: 0xFF05, Hi: 0xFF07, Stride: 0x1},
		{Lo: 0xFF08, Hi: 0xFF08, Stride: 0x1},
		{Lo: 0xFF09, Hi: 0xFF09, Stride: 0x1},
		{Lo: 0xFF0A, Hi: 0xFF0A, Stride: 0x1},
		{Lo: 0xFF0B, Hi: 0xFF0B, Stride: 0x1},
		{Lo: 0xFF0C, Hi: 0xFF0C, Stride: 0x1},
		{Lo: 0xFF0D, Hi: 0xFF0D, Stride: 0x1},
		{Lo: 0xFF0E, Hi: 0xFF0F, Stride: 0x1},
		{Lo: 0xFF10, Hi: 0xFF19, Stride: 0x1},
		{Lo: 0xFF1A, Hi: 0xFF1B, Stride: 0x1},
		{Lo: 0xFF1C, Hi: 0xFF1E, Stride: 0x1},
		{Lo: 0xFF1F, Hi: 0xFF20, Stride: 0x1},
		{Lo: 0xFF21, Hi: 0xFF3A, Stride: 0x1},
		{Lo: 0xFF3B, Hi: 0xFF3B, Stride: 0x1},
		{Lo: 0xFF3C, Hi: 0xFF3C, Stride: 0x1},
		{Lo: 0xFF3D, Hi: 0xFF3D, Stride: 0x1},
		{Lo: 0xFF3E, Hi: 0xFF3E, Stride: 0x1},
		{Lo: 0xFF3F, Hi: 0xFF3F, Stride: 0x1},
		{Lo: 0xFF40, Hi: 0xFF40, Stride: 0x1},
		{Lo: 0xFF41, Hi: 0xFF5A, Stride: 0x1},
		{Lo: 0xFF5B, Hi: 0xFF5B, Stride: 0x1},
		{Lo: 0xFF5C, Hi: 0xFF5C, Stride: 0x1},
		{Lo: 0xFF5D, Hi: 0xFF5D, Stride: 0x1},
		{Lo: 0xFF5E, Hi: 0xFF5E, Stride: 0x1},
		{Lo: 0xFF5F, Hi: 0xFF5F, Stride: 0x1},
		{Lo: 0xFF60, Hi: 0xFF60, Stride: 0x1},
		{Lo: 0xFFE0, Hi: 0xFFE1, Stride: 0x1},
		{Lo: 0xFFE2, Hi: 0xFFE2, Stride: 0x1},
		{Lo: 0xFFE3, Hi: 0xFFE3, Stride: 0x1},
		{Lo: 0xFFE4, Hi: 0xFFE4, Stride: 0x1},
		{Lo: 0xFFE5, Hi: 0xFFE6, Stride: 0x1},
	},
	R32: []unicode.Range32{
		{Lo: 0x16FE0, Hi: 0x16FE0, Stride: 0x1},
		{Lo: 0x17000, Hi: 0x187EC, Stride: 0x1},
		{Lo: 0x18800, Hi: 0x18AF2, Stride: 0x1},
		{Lo: 0x1B000, Hi: 0x1B001, Stride: 0x1},
		{Lo: 0x1F004, Hi: 0x1F004, Stride: 0x1},
		{Lo: 0x1F0CF, Hi: 0x1F0CF, Stride: 0x1},
		{Lo: 0x1F18E, Hi: 0x1F18E, Stride: 0x1},
		{Lo: 0x1F191, Hi: 0x1F19A, Stride: 0x1},
		{Lo: 0x1F200, Hi: 0x1F202, Stride: 0x1},
		{Lo: 0x1F210, Hi: 0x1F23B, Stride: 0x1},
		{Lo: 0x1F240, Hi: 0x1F248, Stride: 0x1},
		{Lo: 0x1F250, Hi: 0x1F251, Stride: 0x1},
		{Lo: 0x1F300, Hi: 0x1F320, Stride: 0x1},
		{Lo: 0x1F32D, Hi: 0x1F335, Stride: 0x1},
		{Lo: 0x1F337, Hi: 0x1F37C, Stride: 0x1},
		{Lo: 0x1F37E, Hi: 0x1F393, Stride: 0x1},
		{Lo: 0x1F3A0, Hi: 0x1F3CA, Stride: 0x1},
		{Lo: 0x1F3CF, Hi: 0x1F3D3, Stride: 0x1},
		{Lo: 0x1F3E0, Hi: 0x1F3F0, Stride: 0x1},
		{Lo: 0x1F3F4, Hi: 0x1F3F4, Stride: 0x1},
		{Lo: 0x1F3F8, Hi: 0x1F3FA, Stride: 0x1},
		{Lo: 0x1F3FB, Hi: 0x1F3FF, Stride: 0x1},
		{Lo: 0x1F400, Hi: 0x1F43E, Stride: 0x1},
		{Lo: 0x1F440, Hi: 0x1F440, Stride: 0x1},
		{Lo: 0x1F442, Hi: 0x1F4FC, Stride: 0x1},
		{Lo: 0x1F4FF, Hi: 0x1F53D, Stride: 0x1},
		{Lo: 0x1F54B, Hi: 0x1F54E, Stride: 0x1},
		{Lo: 0x1F550, Hi: 0x1F567, Stride: 0x1},
		{Lo: 0x1F57A, Hi: 0x1F57A, Stride: 0x1},
		{Lo: 0x1F595, Hi: 0x1F596, Stride: 0x1},
		{Lo: 0x1F5A4, Hi: 0x1F5A4, Stride: 0x1},
		{Lo: 0x1F5FB, Hi: 0x1F5FF, Stride: 0x1},
		{Lo: 0x1F600, Hi: 0x1F64F, Stride: 0x1},
		{Lo: 0x1F680, Hi: 0x1F6C5, Stride: 0x1},
		{Lo: 0x1F6CC, Hi: 0x1F6CC, Stride: 0x1},
		{Lo: 0x1F6D0, Hi: 0x1F6D2, Stride: 0x1},
		{Lo: 0x1F6EB, Hi: 0x1F6EC, Stride: 0x1},
		{Lo: 0x1F6F4, Hi: 0x1F6F6, Stride: 0x1},
		{Lo: 0x1F910, Hi: 0x1F91E, Stride: 0x1},
		{Lo: 0x1F920, Hi: 0x1F927, Stride: 0x1},
		{Lo: 0x1F930, Hi: 0x1F930, Stride: 0x1},
		{Lo: 0x1F933, Hi: 0x1F93E, Stride: 0x1},
		{Lo: 0x1F940, Hi: 0x1F94B, Stride: 0x1},
		{Lo: 0x1F950, Hi: 0x1F95E, Stride: 0x1},
		{Lo: 0x1F980, Hi: 0x1F991, Stride: 0x1},
		{Lo: 0x1F9C0, Hi: 0x1F9C0, Stride: 0x1},
		{Lo: 0x20000, Hi: 0x2A6D6, Stride: 0x1},
		{Lo: 0x2A6D7, Hi: 0x2A6FF, Stride: 0x1},
		{Lo: 0x2A700, Hi: 0x2B734, Stride: 0x1},
		{Lo: 0x2B735, Hi: 0x2B73F, Stride: 0x1},
		{Lo: 0x2B740, Hi: 0x2B81D, Stride: 0x1},
		{Lo: 0x2B81E, Hi: 0x2B81F, Stride: 0x1},
		{Lo: 0x2B820, Hi: 0x2CEA1, Stride: 0x1},
		{Lo: 0x2CEA2, Hi: 0x2F7FF, Stride: 0x1},
		{Lo: 0x2F800, Hi: 0x2FA1D, Stride: 0x1},
		{Lo: 0x2FA1E, Hi: 0x2FFFD, Stride: 0x1},
		{Lo: 0x30000, Hi: 0x3FFFD, Stride: 0x1},
	},
	LatinOffset: 0,
}
