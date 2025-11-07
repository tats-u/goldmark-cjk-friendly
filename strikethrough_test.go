package cjkfriendly_test

import (
	"testing"

	. "github.com/tats-u/goldmark-cjk-friendly/v2"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
)

func TestStrikethrough(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			CJKFriendlyStrikethrough,
		),
	)
	testutil.DoTestCaseFile(markdown, "_test/strikethrough.txt", t, testutil.ParseCliCaseArg()...)
}
