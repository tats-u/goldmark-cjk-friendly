package cjkfriendly

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

type cjkFriendlyEmphasis struct {
}

var CJKFriendlyEmphasis = &cjkFriendlyEmphasis{}

func (e *cjkFriendlyEmphasis) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewCJKFriendlyEmphasisParser(), 499),
	))
}
