package hclparser

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func parseProvider(f *hclwrite.File, bl string) error {

	baf := NewBlockAppendFilter("provider.aws", "default_tags", true)
	f, err := baf.Filter(f)
	if err != nil {
		return err
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.tags.unityname", bl), "default_tags", false)
	return err
}
