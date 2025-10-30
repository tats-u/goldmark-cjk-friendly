package cjkfriendly

import (
	"github.com/tats-u/goldmark-cjk-friendly/internal"
	"github.com/yuin/goldmark"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type strikethroughDelimiterProcessor struct {
	isCJKFriendly bool
}

// IsCJKFriendly implements parser.DelimiterProcessor.
func (p *strikethroughDelimiterProcessor) IsCJKFriendly() bool {
	return p.isCJKFriendly
}

func (p *strikethroughDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '~'
}

func (p *strikethroughDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *strikethroughDelimiterProcessor) OnMatch(consumes int) gast.Node {
	return ast.NewStrikethrough()
}

var defaultStrikethroughDelimiterProcessor = &strikethroughDelimiterProcessor{}

type strikethroughParser struct {
	delimitorProcessor *strikethroughDelimiterProcessor
}

var defaultStrikethroughParser = &strikethroughParser{
	delimitorProcessor: defaultStrikethroughDelimiterProcessor,
}

// NewStrikethroughParser return a new InlineParser that parses
// strikethrough expressions.
func NewCJKFriendlyStrikethroughParser() parser.InlineParser {
	return defaultStrikethroughParser
}

func (s *strikethroughParser) Trigger() []byte {
	return []byte{'~'}
}

func (s *strikethroughParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := internal.ScanDelimiter(line, before, 1, s.delimitorProcessor, internal.TwoPrecedingCharacterReader{Reader: block}.TwoPrecedingCharacter)
	if node == nil || node.OriginalLength > 2 || before == '~' {
		return nil
	}

	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}

func (s *strikethroughParser) CloseBlock(parent gast.Node, pc parser.Context) {
	// nothing to do
}

type cjkFriendlyStrikethrough struct {
}

// CJKFriendlyStrikethrough is an extension that allow you to use strikethrough expression like '~~text~~' .
var CJKFriendlyStrikethrough = &cjkFriendlyStrikethrough{}

func (e *cjkFriendlyStrikethrough) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewCJKFriendlyStrikethroughParser(), 499),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(extension.NewStrikethroughHTMLRenderer(), 500),
	))
}
