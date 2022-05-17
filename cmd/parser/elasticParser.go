package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func addAttribute(f *hclwrite.File, name string, value string, replace bool) (*hclwrite.File, error) {
	if !checkAttribute(f, name) {
		aaf := NewAttributeAppendFilter(name, value, false)
		return aaf.Filter(f)
	} else if replace {
		aaf := NewAttributeSetFilter(name, value)
		return aaf.Filter(f)
	}
	return f, nil
}

func checkAttribute(f *hclwrite.File, name string) bool {
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

func parseElastic(f *hclwrite.File, bl string) error {

	//ES version
	f, err := addAttribute(f, fmt.Sprintf("%v.elasticsearch_version", bl), "7.10", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Tags
	baf := NewBlockAppendFilter(bl, "tags", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.tags.unityname", bl), "myunitytag", false)
	if err != nil {
		return err
	}
	//Cluster Config
	baf = NewBlockAppendFilter(bl, "cluster_config", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.cluster_config.instance_type", bl), "i2.xlarge.elasticsearch", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.cluster_config.instance_count", bl), "2", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	f, err = addAttribute(f, fmt.Sprintf("%v.cluster_config.zone_awareness_enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Test
	//baf = NewBlockAppendFilter(bl+".cluster_config", "zone_awareness_config", true)
	//f, err = baf.Filter(f)
	//if err != nil {
	//	fmt.Printf("%v", err)
	//}
	//f, err = addAttribute(f, fmt.Sprintf("%v.cluster_config.zone_awareness_config.zone_awareness_enabled2", bl), "true", false)
	//if err != nil {
	//	fmt.Printf("%v", err)
	//}

	//VPC Options
	baf = NewBlockAppendFilter(bl, "vpc_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.vpc_options.subnet_ids", bl), "[\n      aws_subnet.subnet-uno.id,\n      aws_subnet.subnet-two.id,\n    ]", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.vpc_options.security_group_ids", bl), "[aws_security_group.es.id]", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//EBS OPTION
	baf = NewBlockAppendFilter(bl, "ebs_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.ebs_options.ebs_enabled", bl), "false", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Advanced Security Options
	baf = NewBlockAppendFilter(bl, "advanced_security_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.advanced_security_options.enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.advanced_security_options.internal_user_database_enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Domain Endpoint Options
	baf = NewBlockAppendFilter(bl, "domain_endpoint_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.domain_endpoint_options.enforce_https", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	f, err = addAttribute(f, fmt.Sprintf("%v.domain_endpoint_options.tls_security_policy", bl), "Policy-Min-TLS-1-2-2019-07", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Node to Node Encryption
	baf = NewBlockAppendFilter(bl, "node_to_node_encryption", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.node_to_node_encryption.enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Encrypt At Rest
	baf = NewBlockAppendFilter(bl, "encrypt_at_rest", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = addAttribute(f, fmt.Sprintf("%v.encrypt_at_rest.enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	return nil
}
