package cjkfriendly

import "github.com/yuin/goldmark"

type cjkFriendlyEmphasisAndStrikethrough struct {
}

// Composite extension of CJKFriendlyEmphasis and CJKFriendlyStrikethrough.
//
// It is recommended to use this extension instead of using CJKFriendlyEmphasis and CJKFriendlyStrikethrough separately.
var CJKFriendlyEmphasisAndStrikethrough = &cjkFriendlyEmphasisAndStrikethrough{}

func (e *cjkFriendlyEmphasisAndStrikethrough) Extend(m goldmark.Markdown) {
	CJKFriendlyEmphasis.Extend(m)
	CJKFriendlyStrikethrough.Extend(m)
}
