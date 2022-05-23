package hclparser

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func parseEC2(f *hclwrite.File, bl string) error {
	//AMI
	f, err := addAttribute(f, fmt.Sprintf("%v.ami", bl), "amiid", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Instance Type
	f, err = addAttribute(f, fmt.Sprintf("%v.instance_type", bl), "t3.xlarge", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Key Name
	f, err = addAttribute(f, fmt.Sprintf("%v.key_name", bl), "", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//TODO
	//VPC Security Ids
	/*f, err = addAttribute(f, fmt.Sprintf("%v.vpc_security_group_ids", bl), "", false)
	if err != nil {
		fmt.Printf("%v", err)
	}*/

	//TODO
	// Security Groups
	f, err = addAttribute(f, fmt.Sprintf("%v.vpc_security_group_ids", bl), "", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// Tags

	return nil
}
