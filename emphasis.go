package cjkfriendly

import (
	"github.com/tats-u/goldmark-cjk-friendly/v2/internal"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type cjkFriendlyEmphasisDelimiterProcessor struct {
}

func (p *cjkFriendlyEmphasisDelimiterProcessor) IsDelimiter(b byte) bool {
	return b == '*' || b == '_'
}

func (p *cjkFriendlyEmphasisDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
	return opener.Char == closer.Char
}

func (p *cjkFriendlyEmphasisDelimiterProcessor) OnMatch(consumes int) ast.Node {
	return ast.NewEmphasis(consumes)
}

var defaultCJKFriendlyEmphasisDelimiterProcessor = &cjkFriendlyEmphasisDelimiterProcessor{}

type cjkFriendlyEmphasisParser struct {
	EmphasisDelimiterProcessor *cjkFriendlyEmphasisDelimiterProcessor
}

var defaultCJKEmphasisParser = &cjkFriendlyEmphasisParser{
	EmphasisDelimiterProcessor: defaultCJKFriendlyEmphasisDelimiterProcessor,
}

// NewCJKFriendlyEmphasisParser return a new InlineParser that parses emphasises.
func NewCJKFriendlyEmphasisParser() parser.InlineParser {
	return defaultCJKEmphasisParser
}

func (s *cjkFriendlyEmphasisParser) Trigger() []byte {
	return []byte{'*', '_'}
}

func (s *cjkFriendlyEmphasisParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	before := block.PrecendingCharacter()
	line, segment := block.PeekLine()
	node := internal.ScanDelimiter(line, before, 1, s.EmphasisDelimiterProcessor, (internal.TwoPrecedingCharacterReader{Reader: block}).TwoPrecedingCharacter)
	if node == nil {
		return nil
	}
	node.Segment = segment.WithStop(segment.Start + node.OriginalLength)
	block.Advance(node.OriginalLength)
	pc.PushDelimiter(node)
	return node
}
