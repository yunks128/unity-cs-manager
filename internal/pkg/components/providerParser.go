package components

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/hclparser"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/tagging"
)

func parseProvider(f *hclwrite.File, bl string, tags tagging.Mandatorytags) error {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(tags)
	_ = json.Unmarshal(inrec, &inInterface)

	var s []string
	for field, val := range inInterface {
		s = append(s, fmt.Sprintf("%v = \"%v\"", field, val))
	}

	baf := hclparser.NewBlockAppendFilter("provider.aws", "default_tags", true)
	f, err := baf.Filter(f)
	if err != nil {
		return err
	}
	_, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.default_tags.tags", "provider.aws"), fmt.Sprintf("{%v}", strings.Join(s, ",")), false)
	return err
}
