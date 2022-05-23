package hclparser

import (
	"errors"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"strings"
)

// AttributeGetSink is a sink implementation for getting a value of attribute.
type AttributeGetSink struct {
	address string
}

var _ Sink = (*AttributeGetSink)(nil)

// NewAttributeGetSink creates a new instance of AttributeGetSink.
func NewAttributeGetSink(address string) Sink {
	return &AttributeGetSink{
		address: address,
	}
}

// Sink reads HCL and writes value of attribute.
func (s *AttributeGetSink) Sink(inFile *hclwrite.File) ([]string, error) {
	attr, _, err := findAttribute(inFile.Body(), s.address)
	if err != nil {
		return nil, err
	}

	// not found
	if attr == nil {
		return []string{}, nil
	}

	// treat expr as a string without interpreting its meaning.
	out, err := GetAttributeValueAsString(attr)

	if err != nil {
		return []string{}, err
	}

	return []string{out + "\n"}, nil
}
func findAttribute(body *hclwrite.Body, address string) (*hclwrite.Attribute, *hclwrite.Body, error) {
	if len(address) == 0 {
		return nil, nil, errors.New("failed to parse address. address is empty")
	}

	a := strings.Split(address, ".")
	if len(a) == 1 {
		// if the address does not cantain any dots, find attribute in the body.
		attr := body.GetAttribute(a[0])
		return attr, body, nil
	}

	// if address contains dots, the last element is an attribute name,
	// and the rest is the address of the block.
	attrName := a[len(a)-1]
	blockAddr := strings.Join(a[:len(a)-1], ".")
	blocks, err := findLongestMatchingBlocks(body, blockAddr)
	if err != nil {
		return nil, nil, err
	}

	if len(blocks) == 0 {
		// not found
		return nil, nil, nil
	}

	// if blocks are matched, check if it has a given attribute name
	for _, b := range blocks {
		attr := b.Body().GetAttribute(attrName)
		if attr != nil {
			// return first matching one.
			return attr, b.Body(), nil
		}
	}

	// not found
	return nil, nil, nil
}

func GetAttributeValueAsString(attr *hclwrite.Attribute) (string, error) {
	// find TokenEqual
	expr := attr.Expr()
	exprTokens := expr.BuildTokens(nil)

	// append tokens until find TokenComment
	var valueTokens hclwrite.Tokens
	for _, t := range exprTokens {
		if t.Type == hclsyntax.TokenComment {
			break
		}
		valueTokens = append(valueTokens, t)
	}

	// TokenIdent records SpaceBefore, but we should ignore it here.
	value := strings.TrimSpace(string(valueTokens.Bytes()))

	return value, nil
}
