package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/hclparser"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/tagging"
)

func Runp(path string, tags tagging.Mandatorytags, subnet, secgroup []string) {
	items, _ := ioutil.ReadDir(path)
	for _, item := range items {
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir(item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					// handle file there
					fmt.Println(item.Name() + "/" + subitem.Name())
					parseFile(item.Name()+"/"+subitem.Name(), tags, subnet, secgroup)
				}
			}
		} else {
			// handle file there
			if strings.HasSuffix(item.Name(), ".tf") {
				fmt.Println(item.Name())
				parseFile(path+"/"+item.Name(), tags, subnet, secgroup)
			}
		}
	}
}

func parseFile(path string, tags tagging.Mandatorytags, subnet, secgroup []string) {
	fbyte, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fwr, diag := hclwrite.ParseConfig(fbyte, "output/"+filepath.Base(path), hcl.Pos{Line: 1, Column: 1})
	if diag.HasErrors() {
		fmt.Println("Couldn't parse file")
	}
	blocks, err := getBlocks(fwr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, b := range blocks {
		fmt.Println(b)
		splits := strings.Split(b, ".")
		blocktype := ""
		if len(splits) > 1 {
			blocktype = splits[1]
		}
		switch blocktype {
		case "aws_elasticsearch_domain":
			// err = parseElastic(fwr, "resource.aws_elasticsearch_domain.unity-sample")
			err = parseElastic(fwr, b, subnet, secgroup, tags.Venue, tags.Project)
			if err != nil {
				fmt.Printf("%v", err)
			}
		case "aws_instance":
			err = parseEC2(fwr, b, subnet, secgroup, tags.Venue)
			if err != nil {
				fmt.Printf("%v", err)
			}
		case "aws":
			err = parseProvider(fwr, blocktype, tags)
			if err != nil {
				fmt.Printf("%v", err)
			}
		case "":
			fmt.Println("three")
		}
	}
	fmt.Printf("%v", string(fwr.Bytes()))
	tfFile, _ := os.Create("output/" + filepath.Base(path))
	_, _ = tfFile.Write(fwr.Bytes())
}

func getBlocks(f *hclwrite.File) ([]string, error) {
	bl := hclparser.NewBlockListSink()
	return bl.Sink(f)
}

/*func addTagsToBlocks(f *hclwrite.File) error {
	blocks, err := getBlocks(f)
	if err != nil {
		return err
	}
	for _, bl := range blocks {
		aaf := hclparser.NewAttributeAppendFilter(fmt.Sprintf("%v.tags.unityname", bl), "myunitydeployment", false)
		f, err = aaf.Filter(f)
		if err != nil {
			return err
		}
	}

	baf := hclparser.NewBlockAppendFilter("resource.aws_eip.ip-test-env", "tags", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf := hclparser.NewAttributeAppendFilter(fmt.Sprintf("%v.tags.unityname", "resource.aws_eip.ip-test-env"), "myunitydeployment", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	/*as := attributeSet{
		address: "resource.aws_eip.ip-test-env.tags.unityname",
		value:   "myunitydeployment",
	}
	f, err = as.Filter(f)
	return err
}*/
