package hclparser

import "github.com/hashicorp/hcl/v2/hclwrite"

func AddAttribute(f *hclwrite.File, name string, value string, replace bool) (*hclwrite.File, error) {
	if !CheckAttribute(f, name) {
		aaf := NewAttributeAppendFilter(name, value, true)
		return aaf.Filter(f)
	} else if replace {
		aaf := NewAttributeSetFilter(name, value)
		return aaf.Filter(f)
	}
	return f, nil
}

func CheckAttribute(f *hclwrite.File, name string) bool {
	aaf := NewAttributeGetSink(name)
	resp, err := aaf.Sink(f)
	if err != nil {
		return false
	}
	if len(resp) > 0 {
		return true
	}
	return false
}
