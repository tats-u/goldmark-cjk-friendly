package internal

import (
	"unicode/utf8"

	"github.com/yuin/goldmark/text"
)

type TwoPrecedingCharacterReader struct {
	Reader text.Reader
}

func (r TwoPrecedingCharacterReader) TwoPrecedingCharacter(beforeLength int) (rune, bool) {
	_, pos := r.Reader.Position()
	source := r.Reader.Source()
	if pos.Start <= beforeLength {
		return utf8.RuneError, false
	}
	rn, size := utf8.DecodeLastRune(source[:(pos.Start - beforeLength)])
	return rn, rn != utf8.RuneError || size != 3
}
