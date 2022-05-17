package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

// BlockAppendFilter is a filter implementation for appending block.
type BlockAppendFilter struct {
	parent  string
	child   string
	newline bool
}

var _ Filter = (*BlockAppendFilter)(nil)

// NewBlockAppendFilter creates a new instance of BlockAppendFilter.
func NewBlockAppendFilter(parent string, child string, newline bool) Filter {
	return &BlockAppendFilter{
		parent:  parent,
		child:   child,
		newline: newline,
	}
}

// Filter reads HCL and appends only matched blocks at a given address.
// The child address is relative to parent one.
// If a newline flag is true, it also appends a newline before the new block.
func (f *BlockAppendFilter) Filter(inFile *hclwrite.File) (*hclwrite.File, error) {
	pTypeName, pLabels, err := parseAddress(f.parent)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parent address: %s", err)
	}

	cTypeName, cLabels, err := parseAddress(f.child)
	if err != nil {
		return nil, fmt.Errorf("failed to parse child address: %s", err)
	}

	matched := findBlocks(inFile.Body(), pTypeName, pLabels)

	for _, b := range matched {
		if f.newline {
			b.Body().AppendNewline()
		}
		b.Body().AppendNewBlock(cTypeName, cLabels)
	}

	return inFile, nil
}

func parseAddress(address string) (string, []string, error) {
	if len(address) == 0 {
		return "", []string{}, fmt.Errorf("failed to parse address: %s", address)
	}

	a := strings.Split(address, ".")
	typeName := a[0]
	labels := []string{}
	if len(a) > 1 {
		labels = a[1:]
	}
	return typeName, labels, nil
}

func findBlocks(b *hclwrite.Body, typeName string, labels []string) []*hclwrite.Block {
	var matched []*hclwrite.Block
	for _, block := range b.Blocks() {
		if typeName == block.Type() {
			if len(labels) > 2 {
				rblock := findNestedBlocks(block.Body().Blocks(), labels[2:])
				if rblock != nil {
					matched = append(matched, rblock)
				}
			}
			labelNames := block.Labels()
			if len(labels) == 0 && len(labelNames) == 0 {
				matched = append(matched, block)
				continue
			}
			if matchLabels(labels, labelNames) {
				matched = append(matched, block)
			}
		}
	}

	return matched
}

func findNestedBlocks(blocks []*hclwrite.Block, labels []string) *hclwrite.Block {
	for i, bl := range blocks {
		if bl.Type() == labels[i] {
			if len(labels) > 1 {
				findNestedBlocks(bl.Body().Blocks(), labels[1:])
			} else {
				return bl
			}
		}
	}
	return nil
}

func matchLabels(lhs []string, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for i := range lhs {
		if !(lhs[i] == rhs[i] || lhs[i] == "*" || rhs[i] == "*") {
			return false
		}
	}

	return true
}
