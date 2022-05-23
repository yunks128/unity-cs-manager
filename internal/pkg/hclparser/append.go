package hclparser

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"log"
	"runtime/debug"
	"strings"
)

// AttributeAppendFilter is a filter implementation for appending attribute.
type AttributeAppendFilter struct {
	address string
	value   string
	newline bool
}

var _ Filter = (*AttributeAppendFilter)(nil)

// NewAttributeAppendFilter creates a new instance of AttributeAppendFilter.
func NewAttributeAppendFilter(address string, value string, newline bool) Filter {
	return &AttributeAppendFilter{
		address: address,
		value:   value,
		newline: newline,
	}
}

// Filter reads HCL and appends a new attribute to a given address.
// If a matched block not found, nothing happens.
// If the given attribute already exists, it returns an error.
// If a newline flag is true, it also appends a newline before the new attribute.
func (f *AttributeAppendFilter) Filter(inFile *hclwrite.File) (*hclwrite.File, error) {
	attrName := f.address
	body := inFile.Body()

	a := strings.Split(f.address, ".")
	if len(a) > 1 {
		// if address contains dots, the last element is an attribute name,
		// and the rest is the address of the block.
		attrName = a[len(a)-1]
		blockAddr := strings.Join(a[:len(a)-1], ".")
		blocks, err := findLongestMatchingBlocks(body, blockAddr)
		if err != nil {
			return nil, err
		}

		if len(blocks) == 0 {
			// not found
			return inFile, nil
		}

		// Use first matching one.
		body = blocks[0].Body()
		if body.GetAttribute(attrName) != nil {
			return nil, fmt.Errorf("attribute already exists: %s", f.address)
		}
	}

	// To delegate expression parsing to the hclwrite parse,
	// We build a new expression and set back to the attribute by tokens.
	expr, err := buildExpression(attrName, f.value)
	if err != nil {
		return nil, err
	}

	if f.newline {
		body.AppendNewline()
	}
	body.SetAttributeRaw(attrName, expr.BuildTokens(nil))

	return inFile, nil
}

func findLongestMatchingBlocks(body *hclwrite.Body, address string) ([]*hclwrite.Block, error) {
	if len(address) == 0 {
		return nil, errors.New("failed to parse address. address is empty")
	}

	a := strings.Split(address, ".")
	typeName := a[0]
	blocks := allMatchingBlocksByType(body, typeName)

	if len(a) == 1 {
		// if the address does not cantain any dots,
		// return all matching blocks by type
		return blocks, nil
	}

	matched := []*hclwrite.Block{}
	// if address contains dots, the next element maybe label or nested block.
	for _, b := range blocks {
		labels := b.Labels()
		// consume labels from address
		matchedlabels := longestMatchingLabels(labels, a[1:])
		if len(matchedlabels) < len(labels) {
			// The labels take precedence over nested blocks.
			// If extra labels remain, skip it.
			continue
		}
		if len(matchedlabels) < (len(a)-1) || len(labels) == 0 {
			// if the block has no labels or partially matched ones, find the nested block
			nestedAddr := strings.Join(a[1+len(matchedlabels):], ".")
			nested, err := findLongestMatchingBlocks(b.Body(), nestedAddr)
			if err != nil {
				return nil, err
			}
			matched = append(matched, nested...)
			continue
		}
		// all labels are matched, just add it to matched list.
		matched = append(matched, b)
	}

	return matched, nil
}

func allMatchingBlocksByType(b *hclwrite.Body, typeName string) []*hclwrite.Block {
	matched := []*hclwrite.Block{}
	for _, block := range b.Blocks() {
		if typeName == block.Type() {
			matched = append(matched, block)
		}
	}

	return matched
}

func longestMatchingLabels(labels []string, prefix []string) []string {
	matched := []string{}
	for i := range prefix {
		if len(labels) <= i {
			return matched
		}
		if prefix[i] != labels[i] {
			return matched
		}
		matched = append(matched, labels[i])
	}
	return matched
}

func safeParseConfig(src []byte, filename string, start hcl.Pos) (f *hclwrite.File, e error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[DEBUG] failed to parse input: %s\nstacktrace: %s", filename, string(debug.Stack()))
			// Set a return value from panic recover
			e = fmt.Errorf(`failed to parse input: %s
panic: %s
This may be caused by a bug in the hclwrite parse`, filename, err)
		}
	}()

	f, diags := hclwrite.ParseConfig(src, filename, start)

	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse input: %s", diags)
	}

	return f, nil
}
