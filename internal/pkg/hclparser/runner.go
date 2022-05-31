package hclparser

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"io/ioutil"
	"log"
	"strings"
)

type networkinfo struct {
	subnetid         []string
	securitygroupids []string
}

type tag struct {
	tags map[string]string
}

func Runp(path string, tags, subnet, secgroup []string) {

	fmt.Printf("hello")
	items, _ := ioutil.ReadDir(path)
	for _, item := range items {
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir(item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					// handle file there
					fmt.Println(item.Name() + "/" + subitem.Name())
					parseFile(item.Name() + "/" + subitem.Name())
				}
			}
		} else {
			// handle file there
			fmt.Println(item.Name())
		}
	}
}

func parseFile(path string) {
	fbyte, err := ioutil.ReadFile("test/elastic.tf")
	if err != nil {
		fmt.Println(err)
	}
	fwr, diag := hclwrite.ParseConfig(fbyte, "example.tf", hcl.Pos{Line: 1, Column: 1})
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
			//err = parseElastic(fwr, "resource.aws_elasticsearch_domain.unity-sample")
			err = parseElastic(fwr, blocktype)
			if err != nil {
				fmt.Printf("%v", err)
			}
		case "aws_instance":
			err = parseEC2(fwr, blocktype)
			if err != nil {
				fmt.Printf("%v", err)
			}
		case "":
			fmt.Println("three")
		}
	}
	//func main() {
	//
	//	fbyte, err := ioutil.ReadFile("test/elastic.tf")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fwr, diag := hclwrite.ParseConfig(fbyte, "example.tf", hcl.Pos{Line: 1, Column: 1})
	//	if diag.HasErrors() {
	//		fmt.Println("Couldn't parse file")
	//	}
	//	//fmt.Printf("%v", string(fwr.Bytes()))
	//	//aaf := NewAttributeAppendFilter("resource.aws_instance.unity-ec2-instance.tags.unityname", "mytag", false)
	//	//fwr, err = aaf.Filter(fwr)
	//	//err = addTagsToBlocks(fwr)
	//	blocks, err := getBlocks(fwr)
	//	if err != nil {
	//		log.Fatalf("%v", err)
	//	}
	//	for _, b := range blocks {
	//		fmt.Println(b)
	//		splits := strings.Split(b, ".")
	//		blocktype := ""
	//		if len(splits) > 1 {
	//			blocktype = splits[1]
	//		}
	//		switch blocktype {
	//		case "aws_elasticsearch_domain":
	//			//err = parseElastic(fwr, "resource.aws_elasticsearch_domain.unity-sample")
	//			err = parseElastic(fwr, blocktype)
	//			if err != nil {
	//				fmt.Printf("%v", err)
	//			}
	//		case "aws_instance":
	//			err = parseEC2(fwr, blocktype)
	//			if err != nil {
	//				fmt.Printf("%v", err)
	//			}
	//		case "":
	//			fmt.Println("three")
	//		}
	//	}
	//	fmt.Printf("%v", string(fwr.Bytes()))
	//
	//}
}

func getBlocks(f *hclwrite.File) ([]string, error) {

	bl := NewBlockListSink()
	return bl.Sink(f)

}

func addTagsToBlocks(f *hclwrite.File) error {
	blocks, err := getBlocks(f)
	if err != nil {
		return err
	}
	for _, bl := range blocks {
		aaf := NewAttributeAppendFilter(fmt.Sprintf("%v.tags.unityname", bl), "myunitydeployment", false)
		f, err = aaf.Filter(f)
		if err != nil {
			return err
		}
	}

	baf := NewBlockAppendFilter("resource.aws_eip.ip-test-env", "tags", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf := NewAttributeAppendFilter(fmt.Sprintf("%v.tags.unityname", "resource.aws_eip.ip-test-env"), "myunitydeployment", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	/*as := attributeSet{
		address: "resource.aws_eip.ip-test-env.tags.unityname",
		value:   "myunitydeployment",
	}
	f, err = as.Filter(f)*/
	return err
}
