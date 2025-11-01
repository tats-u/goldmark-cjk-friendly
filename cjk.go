package cjkfriendly

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

type cjkFriendlyEmphasis struct {
}

// CJKFriendlyEmphasis is a basic extension without GFM strikethrough support
//
// Use CJKFriendlyEmphasisAndStrikethrough instead if you need GFM strikethrough support
var CJKFriendlyEmphasis = &cjkFriendlyEmphasis{}

func (e *cjkFriendlyEmphasis) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewCJKFriendlyEmphasisParser(), 499),
	))
}
