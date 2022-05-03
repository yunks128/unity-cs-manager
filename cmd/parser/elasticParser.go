package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func parseElastic(f *hclwrite.File) error {

	bl := "resource.aws_elasticsearch_domain.unity-sample"
	//ES version
	aaf := NewAttributeAppendFilter(fmt.Sprintf("%v.elasticsearch_version", bl), "7.10", false)
	f, err := aaf.Filter(f)
	if err != nil {
		return err
	}

	//Tags
	baf := NewBlockAppendFilter(bl, "tags", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.tags.unityname", bl), "myunitytag", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	//Cluster Config

	//VPC Options

	//EBS OPTION

	//Advanced Security Options

	//Domain Endpoint Options

	//Node to Node Encryption

	//Encrypt At Rest

	return nil
}
