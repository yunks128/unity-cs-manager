package hclparser

import "github.com/hashicorp/hcl/v2/hclwrite"

type Sink interface {
	// Sink reads HCL and writes bytes.
	Sink(*hclwrite.File) ([]string, error)
}

type Filter interface {
	// Filter reads HCL and writes HCL
	Filter(*hclwrite.File) (*hclwrite.File, error)
}
