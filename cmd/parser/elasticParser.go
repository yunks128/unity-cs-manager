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
	baf = NewBlockAppendFilter(bl, "cluster_config", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.cluster_config.instance_type", bl), "i2.xlarge.elasticsearch", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.cluster_config.instance_count", bl), "2", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.cluster_config.zone_awareness_enabled", bl), "true", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	baf = NewBlockAppendFilter(bl+".cluster_config", "zone_awareness_config", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}

	//VPC Options
	baf = NewBlockAppendFilter(bl, "vpc_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.vpc_options.subnet_ids", bl), "[\n      aws_subnet.subnet-uno.id,\n      aws_subnet.subnet-two.id,\n    ]", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.vpc_options.security_group_ids", bl), "[aws_security_group.es.id]", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}

	//EBS OPTION
	baf = NewBlockAppendFilter(bl, "ebs_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.ebs_options.ebs_enabled", bl), "false", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}

	//Advanced Security Options
	baf = NewBlockAppendFilter(bl, "advanced_security_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.advanced_security_options.enabled", bl), "true", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.advanced_security_options.internal_user_database_enabled", bl), "true", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}

	//Domain Endpoint Options
	baf = NewBlockAppendFilter(bl, "domain_endpoint_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.domain_endpoint_options.enforce_https", bl), "true", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.domain_endpoint_options.tls_security_policy", bl), "Policy-Min-TLS-1-2-2019-07", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}

	//Node to Node Encryption
	baf = NewBlockAppendFilter(bl, "node_to_node_encryption", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.node_to_node_encryption.enabled", bl), "true", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}

	//Encrypt At Rest
	baf = NewBlockAppendFilter(bl, "encrypt_at_rest", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	aaf = NewAttributeAppendFilter(fmt.Sprintf("%v.encrypt_at_rest.enabled", bl), "true", false)
	f, err = aaf.Filter(f)
	if err != nil {
		return err
	}

	return nil
}
