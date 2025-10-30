package cjkfriendly_test

import (
	"testing"

	. "github.com/tats-u/goldmark-cjk-friendly"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/testutil"
)

func TestCJKFriendlyEmphasis(t *testing.T) {
	markdown := goldmark.New(goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	),
		goldmark.WithExtensions(CJKFriendlyEmphasis),
	)
	no := 1
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Only preceding character of the opening mark is a normal CJK character",
			Markdown:    "この**`code`**",
			Expected:    "<p>この<strong><code>code</code></strong></p>",
		},
		t,
	)
	no = 2
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Only following character of the opening mark is a normal CJK character",
			Markdown:    "John**「ハロー」**",
			Expected:    "<p>John<strong>「ハロー」</strong></p>",
		},
		t,
	)
	no = 3
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Only following character of the closing mark is a normal CJK character",
			Markdown:    "**`code`**を",
			Expected:    "<p><strong><code>code</code></strong>を</p>",
		},
		t,
	)
	no = 4
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Only preceding character of the closing mark is a normal CJK character",
			Markdown:    "Git **（ギット）**Hub",
			Expected:    "<p>Git <strong>（ギット）</strong>Hub</p>",
		},
		t,
	)
	no = 5
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes preceding IVS",
			Markdown:    "禰󠄀**(ね)**豆子",
			Expected:    "<p>禰󠄀<strong>(ね)</strong>豆子</p>",
		},
		t,
	)
	no = 6
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the surrounding normal hangul character",
			Markdown:    "**스크립트(script)**라고",
			Expected:    "<p><strong>스크립트(script)</strong>라고</p>",
		},
		t,
	)
	no = 7
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the preceding hangul character that cannot be determined by East Asian Width",
			Markdown:    "ᅡ**(a)**",
			Expected:    "<p>ᅡ<strong>(a)</strong></p>",
		},
		t,
	)
	no = 8
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the following hangul character that cannot be determined by East Asian Width",
			Markdown:    "**(k)**ᄏ",
			Expected:    "<p><strong>(k)</strong>ᄏ</p>",
		},
		t,
	)
	no = 9
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the preceding han SVS",
			Markdown:    "大塚︀**(U+585A U+FE00)** 大塚**(U+FA10)**",
			Expected:    "<p>大塚︀<strong>(U+585A U+FE00)</strong> 大塚<strong>(U+FA10)</strong></p>",
		},
		t,
	)
	no = 10
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the preceding pseudo-emoji CJK symbol",
			Markdown:    "〽︎**(庵点)**は、",
			Expected:    "<p>〽︎<strong>(庵点)</strong>は、</p>",
		},
		t,
	)
	no = 11
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the preceding CJK ambiguous punctuation sequence (regression test)",
			Markdown:    "**“︁Git”︁**Hub",
			Expected:    "<p><strong>“︁Git”︁</strong>Hub</p>",
		},
		t,
	)
	no = 12
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the preceding CJK ambiguous punctuation sequence (underscore)",
			Markdown:    "“︁Git”︁__Hub__",
			Expected:    "<p>“︁Git”︁<strong>Hub</strong></p>",
		},
		t,
	)
}

func TestCJKFriendlyStrikethrough(t *testing.T) {
	markdown := goldmark.New(goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	),
		goldmark.WithExtensions(CJKFriendlyStrikethrough),
	)
	no := 1
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Strikethrough is enabled",
			Markdown:    "~~No~~Yes",
			Expected:    "<p><del>No</del>Yes</p>",
		},
		t,
	)
	no = 2
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the preceding supplementary han character",
			Markdown:    "𩸽~~()a~~a",
			Expected:    "<p>𩸽<del>()a</del>a</p>",
		},
		t,
	)
	no = 3
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the following supplementary han character",
			Markdown:    "a~~a()~~𩸽",
			Expected:    "<p>a<del>a()</del>𩸽</p>",
		},
		t,
	)
}

func TestCJKFriendlyEmphasisWithEscapedSpace(t *testing.T) {
	markdown := goldmark.New(goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	),
		goldmark.WithExtensions(extension.NewCJK(extension.WithEscapedSpace()), CJKFriendlyEmphasis),
	)
	no := 1
	testutil.DoTestCase(
		markdown,
		testutil.MarkdownTestCase{
			No:          no,
			Description: "Recognizes the following supplementary han character",
			Markdown:    "a\\ **()**\\ a𩸽**()**𩸽",
			Expected:    "<p>a<strong>()</strong>a𩸽<strong>()</strong>𩸽</p>",
		},
		t,
	)
}
