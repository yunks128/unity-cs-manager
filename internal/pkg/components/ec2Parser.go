package components

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/hclparser"
	"strings"
)

func parseEC2(f *hclwrite.File, bl string, subnet, secgroup []string, venue string) error {
	//AMI
	f, err := hclparser.AddAttribute(f, fmt.Sprintf("%v.ami", bl), "amiid", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Instance Type
	instancetype := ""
	if venue == "dev" {
		instancetype = "t3.medium"
	} else if venue == "stage" {
		instancetype = "t3.large"
	} else if venue == "prod" {
		instancetype = "t3.xlarge"
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.instance_type", bl), instancetype, false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Key Name
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.key_name", bl), "mykey", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//VPC Security Ids
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.vpc_security_group_ids", bl), fmt.Sprintf("[%v]", strings.Join(secgroup, ",")), false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// Subnet Id
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.subnet_id", bl), fmt.Sprintf("%v\n", subnet[0]), false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// Tags

	return nil
}
