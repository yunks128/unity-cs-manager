package components

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/sethvargo/go-password/password"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/hclparser"
	"strings"
)

func parseElastic(f *hclwrite.File, bl string, subnet []string, securitygroup []string, venue string, project string) error {

	instancesize := ""
	instancecount := ""
	if venue == "dev" {
		instancesize = "i3.large.search"
		instancecount = "1"
	} else if venue == "stage" {
		instancesize = "i3.xlarge.search"
		instancecount = "2"
	} else if venue == "prod" {
		instancesize = "i3.2xlarge.search"
		instancecount = "3"
	}

	//ES version
	f, err := hclparser.AddAttribute(f, fmt.Sprintf("%v.elasticsearch_version", bl), "7.10", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Tags
	baf := hclparser.NewBlockAppendFilter(bl, "tags", true)
	f, err = baf.Filter(f)
	if err != nil {
		return err
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.tags.unityname", bl), "myunitytag", false)
	if err != nil {
		return err
	}
	//Cluster Config
	baf = hclparser.NewBlockAppendFilter(bl, "cluster_config", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.cluster_config.instance_type", bl), instancesize, false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.cluster_config.instance_count", bl), instancecount, false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.cluster_config.zone_awareness_enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//VPC Options
	baf = hclparser.NewBlockAppendFilter(bl, "vpc_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}

	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.vpc_options.subnet_ids", bl), fmt.Sprintf("[%v]", strings.Join(subnet, ",")), false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.vpc_options.security_group_ids", bl), fmt.Sprintf("[%v]", strings.Join(securitygroup, ",")), false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//EBS OPTION
	baf = hclparser.NewBlockAppendFilter(bl, "ebs_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.ebs_options.ebs_enabled", bl), "false", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Advanced Security Options
	baf = hclparser.NewBlockAppendFilter(bl, "advanced_security_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.advanced_security_options.enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.advanced_security_options.internal_user_database_enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	baf = hclparser.NewBlockAppendFilter(bl, "master_user_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.master_user_options.master_user_name", bl), "admin-"+project+"-"+venue, false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	res, err := password.Generate(12, 10, 0, false, false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.master_user_options.master_user_password", bl), fmt.Sprintf("%v", res), false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Domain Endpoint Options
	baf = hclparser.NewBlockAppendFilter(bl, "domain_endpoint_options", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.domain_endpoint_options.enforce_https", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.domain_endpoint_options.tls_security_policy", bl), "Policy-Min-TLS-1-2-2019-07", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Node to Node Encryption
	baf = hclparser.NewBlockAppendFilter(bl, "node_to_node_encryption", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.node_to_node_encryption.enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//Encrypt At Rest
	baf = hclparser.NewBlockAppendFilter(bl, "encrypt_at_rest", true)
	f, err = baf.Filter(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	f, err = hclparser.AddAttribute(f, fmt.Sprintf("%v.encrypt_at_rest.enabled", bl), "true", false)
	if err != nil {
		fmt.Printf("%v", err)
	}

	return nil
}
