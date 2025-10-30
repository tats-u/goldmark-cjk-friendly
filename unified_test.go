package cjkfriendly_test

import (
	"testing"

	. "github.com/tats-u/goldmark-cjk-friendly"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
)

func TestCJKFriendlyEmphasisWithGFM(t *testing.T) {
	markdown := goldmark.New(goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	),
		goldmark.WithExtensions(CJKFriendlyEmphasisAndStrikethrough),
	)
	no := 1
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Only preceding character of the opening mark is a normal CJK character",
			Markdown:    "この~~`wrong code`もとい~~**`correct code`**を太字化",
			Expected:    "<p>この<del><code>wrong code</code>もとい</del><strong><code>correct code</code></strong>を太字化</p>",
		},
		t,
	)
}

func TestCJKFriendlyEmphasisWithEscapedSpaceAndGFM(t *testing.T) {
	markdown := goldmark.New(goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	),
		goldmark.WithExtensions(extension.NewCJK(extension.WithEscapedSpace()), CJKFriendlyEmphasisAndStrikethrough),
	)
	no := 1
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the following supplementary han character",
			Markdown:    "a\\ **()**\\ a𩸽**()**𩸽𠮷~~()~~𠮷",
			Expected:    "<p>a<strong>()</strong>a𩸽<strong>()</strong>𩸽𠮷<del>()</del>𠮷</p>",
		},
		t,
	)
}
