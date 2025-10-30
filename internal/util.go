package internal

import (
	"unicode/utf8"
)

func PrecedingToRune(source []byte, base int) (rune, bool) {
	if base <= 0 {
		return utf8.RuneError, false
	}
	r, size := utf8.DecodeLastRune(source[:base])
	if r == utf8.RuneError && size != 3 {
		return r, false
	}
	return r, true
}

// `true` if the given rune represents a CJK character
//
// Definition: https://github.com/tats-u/markdown-cjk-friendly/blob/main/specification.md#cjk-character
//
// Ranges: https://github.com/tats-u/markdown-cjk-friendly/blob/main/ranges.md#cjk-characters
func IsCJK(r rune) bool {
	// Fast path for European languages including ASCII
	if r < 0x1100 {
		return false
	}
	// Snapshot of Unicode 17
	return /* 0x1100 <= r && */ r <= 0x11ff ||
		r == 0x20a9 ||
		0x2329 <= r && r <= 0x232a ||
		0x2630 <= r && r <= 0x2637 ||
		0x268a <= r && r <= 0x268f ||
		0x2e80 <= r && r <= 0x2e99 ||
		0x2e9b <= r && r <= 0x2ef3 ||
		0x2f00 <= r && r <= 0x2fd5 ||
		0x2ff0 <= r && r <= 0x303e ||
		0x3041 <= r && r <= 0x3096 ||
		0x3099 <= r && r <= 0x30ff ||
		0x3105 <= r && r <= 0x312f ||
		0x3131 <= r && r <= 0x318e ||
		0x3190 <= r && r <= 0x31e5 ||
		0x31ef <= r && r <= 0x321e ||
		0x3220 <= r && r <= 0x3247 ||
		0x3250 <= r && r <= 0xa48c ||
		0xa490 <= r && r <= 0xa4c6 ||
		0xa960 <= r && r <= 0xa97c ||
		0xac00 <= r && r <= 0xd7a3 ||
		0xd7b0 <= r && r <= 0xd7c6 ||
		0xd7cb <= r && r <= 0xd7fb ||
		0xf900 <= r && r <= 0xfaff ||
		0xfe10 <= r && r <= 0xfe19 ||
		0xfe30 <= r && r <= 0xfe52 ||
		0xfe54 <= r && r <= 0xfe66 ||
		0xfe68 <= r && r <= 0xfe6b ||
		0xff01 <= r && r <= 0xffbe ||
		0xffc2 <= r && r <= 0xffc7 ||
		0xffca <= r && r <= 0xffcf ||
		0xffd2 <= r && r <= 0xffd7 ||
		0xffda <= r && r <= 0xffdc ||
		0xffe0 <= r && r <= 0xffe6 ||
		0xffe8 <= r && r <= 0xffee ||
		0x16fe0 <= r && r <= 0x16fe4 ||
		0x16ff0 <= r && r <= 0x16ff6 ||
		0x17000 <= r && r <= 0x18cd5 ||
		0x18cff <= r && r <= 0x18d1e ||
		0x18d80 <= r && r <= 0x18df2 ||
		0x1aff0 <= r && r <= 0x1aff3 ||
		0x1aff5 <= r && r <= 0x1affb ||
		0x1affd <= r && r <= 0x1affe ||
		0x1b000 <= r && r <= 0x1b122 ||
		r == 0x1b132 ||
		0x1b150 <= r && r <= 0x1b152 ||
		r == 0x1b155 ||
		0x1b164 <= r && r <= 0x1b167 ||
		0x1b170 <= r && r <= 0x1b2fb ||
		0x1d300 <= r && r <= 0x1d356 ||
		0x1d360 <= r && r <= 0x1d376 ||
		r == 0x1f200 ||
		r == 0x1f202 ||
		0x1f210 <= r && r <= 0x1f219 ||
		0x1f21b <= r && r <= 0x1f22e ||
		0x1f230 <= r && r <= 0x1f231 ||
		r == 0x1f237 ||
		r == 0x1f23b ||
		0x1f240 <= r && r <= 0x1f248 ||
		0x1f260 <= r && r <= 0x1f265 ||
		0x20000 <= r && r <= 0x3fffd
}

// `true` if the given rune is a non-emoji general purpose variation selector
//
// Definition: https://github.com/tats-u/markdown-cjk-friendly/blob/main/specification.md#non-emoji-general-use-variation-selector
//
// Range: https://github.com/tats-u/markdown-cjk-friendly/blob/main/ranges.md#non-emoji-general-use-variation-selectors
func IsNonEmojiGeneralPurposeVS(r rune) bool {
	return 0xfe00 <= r && r <= 0xfe0e
}

// `true` if the given rune is an ideographic variation selector
//
// https://github.com/tats-u/markdown-cjk-friendly/blob/main/specification.md#ideographic-variation-selector
func IsIdeographicVS(r rune) bool {
	return 0xe0100 <= r && r <= 0xe01ef
}

// `true` if the given runes compose a CJK ambiguous punctuation sequence
//
// Definition: https://github.com/tats-u/markdown-cjk-friendly/blob/main/specification.md#cjk-ambiguous-punctuation-sequence
//
// List: https://github.com/tats-u/markdown-cjk-friendly/blob/main/ranges.md#cjk-ambiguous-punctuation-sequences
func IsCJKAmbiguousPunctuation(base rune, vs rune) bool {
	if vs != 0xfe01 {
		return false
	}
	return base == 0x2018 || base == 0x2019 || base == 0x201c || base == 0x201d
}
