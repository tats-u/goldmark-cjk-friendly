package internal

import (
	"unicode/utf8"

	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

// ScanDelimiter scans a delimiter by given DelimiterProcessor.
func ScanDelimiter(line []byte, before rune, minimum int, processor parser.DelimiterProcessor, getTwoBeforeFromBeforeSize func(int) (rune, bool)) *parser.Delimiter {
	i := 0
	c := line[i]
	j := i
	if !processor.IsDelimiter(c) {
		return nil
	}
	for ; j < len(line) && c == line[j]; j++ {
	}
	if (j - i) >= minimum {
		after := rune(' ')
		if j != len(line) {
			after = util.ToRune(line, j)
		}

		var canOpen, canClose bool
		beforeIsPunctuation := util.IsPunctRune(before)
		beforeIsWhitespace := util.IsSpaceRune(before)
		afterIsPunctuation := util.IsPunctRune(after)
		afterIsWhitespace := util.IsSpaceRune(after)

		var isLeft bool
		var isRight bool

		// Start CJK friendly emphasis process
		isPreviousCJKAmbiguousPunct := false
		hasTwoBefore := false
		if IsNonEmojiGeneralPurposeVS(before) {
			beforeBytes := utf8.RuneLen(before)
			if beforeBytes >= 1 {
				var twoBefore rune
				if twoBefore, hasTwoBefore = getTwoBeforeFromBeforeSize(beforeBytes); hasTwoBefore {
					isPreviousCJKAmbiguousPunct = IsCJKAmbiguousPunctuation(twoBefore, before)
					beforeIsPunctuation = util.IsPunctRune(twoBefore)
					before = twoBefore
				}
			}
		}
		beforeIsCJK := isPreviousCJKAmbiguousPunct || IsCJK(before) || (!hasTwoBefore && IsIdeographicVS(before))
		afterIsCJK := IsCJK(after)
		beforeIsNonCJKPunctuation := beforeIsPunctuation && !beforeIsCJK
		afterIsNonCJKPunctuation := afterIsPunctuation && !afterIsCJK
		isLeft = !afterIsWhitespace &&
			(!afterIsNonCJKPunctuation || beforeIsWhitespace || beforeIsCJK || beforeIsNonCJKPunctuation)
		isRight = !beforeIsWhitespace &&
			(!beforeIsNonCJKPunctuation || afterIsWhitespace || afterIsCJK || afterIsNonCJKPunctuation)
		// End CJK friendly emphasis process

		if line[i] == '_' {
			canOpen = isLeft && (!isRight || beforeIsPunctuation)
			canClose = isRight && (!isLeft || afterIsPunctuation)
		} else {
			canOpen = isLeft
			canClose = isRight
		}
		return parser.NewDelimiter(canOpen, canClose, j-i, c, processor)
	}
	return nil
}
