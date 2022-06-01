package hclparser

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/tagging"
	"strings"
)

func parseProvider(f *hclwrite.File, bl string, tags tagging.Mandatorytags) error {

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(tags)
	json.Unmarshal(inrec, &inInterface)

	var s []string
	for field, val := range inInterface {
		s = append(s, fmt.Sprintf("%v = \"%v\"", field, val))
	}

	baf := NewBlockAppendFilter("provider.aws", "default_tags", true)
	f, err := baf.Filter(f)
	if err != nil {
		return err
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.default_tags.tags", "provider.aws"), fmt.Sprintf("{%v}", strings.Join(s, ",")), false)
	return err
}
