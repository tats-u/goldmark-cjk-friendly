package cjkfriendly

import "github.com/yuin/goldmark"

type cjkFriendlyEmphasisAndStrikethrough struct {
}

var CJKFriendlyEmphasisAndStrikethrough = &cjkFriendlyEmphasisAndStrikethrough{}

func (e *cjkFriendlyEmphasisAndStrikethrough) Extend(m goldmark.Markdown) {
	CJKFriendlyEmphasis.Extend(m)
	CJKFriendlyStrikethrough.Extend(m)
}
