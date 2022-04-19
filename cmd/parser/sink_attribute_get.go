package main

import (
	"errors"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"strings"
)

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
