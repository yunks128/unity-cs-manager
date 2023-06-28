package eks

import (
	"github.com/unity-sds/unity-cs-manager/marketplace"
	"os"

	"bytes"
	"fmt"
	"github.com/unity-sds/unity-cs-manager/internal/pkg/templates"
)
import "html/template"

// Type Name: EKSConfig
// Type Description:
// EKSConfig is a struct type that represents the configuration for an EKS cluster and its associated resources. It contains fields that define the various configuration parameters required to create an EKS cluster and its associated resources.
// The fields of the EKSConfig struct are:
// 1. ServiceArn string: The Amazon Resource Name (ARN) of the service that the EKS cluster uses.
// 2. ClusterName string: The name of the EKS cluster.
// 3. ClusterRegion string: The AWS region in which the EKS cluster is created.
// 4. ClusterVersion string: The version of Kubernetes that the EKS cluster uses.
// 5. ClusterMinSize int: The minimum number of worker nodes in the EKS cluster.
// 6. ClusterMaxSize int: The maximum number of worker nodes in the EKS cluster.
// 7. ClusterDesiredCapacity int: The desired capacity of worker nodes in the EKS cluster.
// 8. ClusterAMI string: The Amazon Machine Image (AMI) ID of the worker nodes in the EKS cluster.
// 9. InstanceRoleArn string: The IAM role ARN that the worker nodes in the EKS cluster use.
// 10. KubeProxyVersion string: The version of kube-proxy that the worker nodes in the EKS cluster use.
// 11. EBSCSIVersion string: The version of the Amazon EBS Container Storage Interface (CSI) driver that the worker nodes in the EKS cluster use.
// 12. CoreDNSVersion string: The version of CoreDNS that the EKS cluster uses.
// 13. PublicSubnetA string: The ID of the first public subnet for the EKS cluster.
// 14. PublicSubnetB string: The ID of the second public subnet for the EKS cluster.
// 15. PrivateSubnetA string: The ID of the first private subnet for the EKS cluster.
// 16. PrivateSubnetB string: The ID of the second private subnet for the EKS cluster.
// 17. SecurityGroup string: The ID of the security group for the EKS cluster.
// 18. SharedNodeSecurityGroup string: The ID of the security group that is shared by the worker nodes in the EKS cluster.
// 19. ClusterInstanceType string: The instance type of the worker nodes in the EKS cluster.
// 20. ClusterOwner string: The owner of the EKS cluster.
// 21. ManagedNodeGroups []NodeGroup: An array of NodeGroup structs that represent the managed node groups of the EKS cluster.
// 22. ServiceName string: The name of the service associated with the EKS cluster.
// 23. ProjectName string: The name of the project associated with the EKS cluster.
// 24. Tags AWSTags: A struct that represents the tags associated with the EKS cluster.
type EKSConfig struct {
	ServiceArn              string
	ClusterName             string
	ClusterRegion           string
	ClusterVersion          string
	ClusterMinSize          int
	ClusterMaxSize          int
	ClusterDesiredCapacity  int
	ClusterAMI              string
	InstanceRoleArn         string
	KubeProxyVersion        string
	EBSCSIVersion           string
	CoreDNSVersion          string
	PublicSubnetA           string
	PublicSubnetB           string
	PrivateSubnetA          string
	PrivateSubnetB          string
	SecurityGroup           string
	SharedNodeSecurityGroup string
	ClusterInstanceType     string
	ClusterOwner            string
	ManagedNodeGroups       []NodeGroup
	ServiceName             string
	ProjectName             string
	Tags                    AWSTags
}

// Type Name: NodeGroup
// Type Description:
// NodeGroup is a struct type that represents a managed node group in an EKS cluster. It contains fields that define the various configuration parameters required to create a managed node group in an EKS cluster.
// The fields of the NodeGroup struct are:
// 1. NodeGroupName string: The name of the managed node group.
// 2. ClusterMinSize int: The minimum number of nodes in the managed node group.
// 3. ClusterMaxSize int: The maximum number of nodes in the managed node group.
// 4. ClusterDesiredCapacity int: The desired capacity of nodes in the managed node group.
// 5. ClusterInstanceType string: The instance type of nodes in the managed node group.
// Example Usage:
//
//	nodeGroup := NodeGroup{
//	    NodeGroupName: "my-node-group",
//	    ClusterMinSize: 1,
//	    ClusterMaxSize: 5,
//	    ClusterDesiredCapacity: 3,
//	    ClusterInstanceType: "t3.medium",
//	}
//
// The above example shows the usage of the NodeGroup struct to define a managed node group in an EKS cluster. The struct contains fields that define the various configuration parameters required to create the managed node group, such as the minimum and maximum number of nodes, the desired capacity, and the instance type of the nodes.
type NodeGroup struct {
	NodeGroupName          string
	ClusterMinSize         int
	ClusterMaxSize         int
	ClusterDesiredCapacity int
	ClusterInstanceType    string
}

// Type Name: AWSTags
// Type Description:
// AWSTags is a struct type that represents the various tags associated with an AWS resource. It contains fields that represent different aspects of a resource such as its name, creator email, project name, service name, etc.
// The fields of the AWSTags struct are:
// 1. Resourcename string: The name of the AWS resource.
// 2. Creatoremail string: The email address of the creator of the AWS resource.
// 3. Pocemail string: The email address of the point of contact (POC) for the AWS resource.
// 4. Venue string: The venue in which the AWS resource is deployed.
// 5. Projectname string: The name of the project associated with the AWS resource.
// 6. Servicename string: The name of the service associated with the AWS resource.
// 7. Applicationname string: The name of the application associated with the AWS resource.
// 8. Applicationversion string: The version of the application associated with the AWS resource.
// 9. Releaseversion string: The version of the AWS resource.
// 10. Componentname string: The name of the component associated with the AWS resource.
// 11. Securityplanid string: The ID of the security plan associated with the AWS resource.
// 12. Exposedweb string: Indicates whether the AWS resource is exposed to the web or not.
// 13. Experimental string: Indicates whether the AWS resource is experimental or not.
// 14. Userfacing string: Indicates whether the AWS resource is user-facing or not.
// 15. Criticalinfra string: Indicates whether the AWS resource is part of the critical infrastructure or not.
// 16. Sourcecontrol string: Indicates the source control system used for the AWS resource.
// Example Usage:
// tags := AWSTags{
//     Resourcename: "my-resource",
//     Creatoremail: "john.doe@example.com",
//     Pocemail: "jane.doe@example.com",
//     Venue: "us-west-2",
//     Projectname: "my-project",
//     Servicename: "my-service",
//     Applicationname: "my-application",
//     Applicationversion: "1.0",
//     Releaseversion: "1.2.3",
//     Componentname: "my-component",
//     Securityplanid: "sp-123456",
//     Exposedweb: "false",
//     Experimental: "true",
//     Userfacing: "true",
//     Criticalinfra: "false",
//     Sourcecontrol: "github",
// }

// The above example shows the usage of the AWSTags struct to define various tags associated with an AWS resource. The struct contains fields that represent different aspects of the resource, such as its name, creator email, project name, service name, etc. These tags can be used to identify and categorize AWS resources based on their various attributes.
type AWSTags struct {
	Resourcename       string
	Creatoremail       string
	Pocemail           string
	Venue              string
	Projectname        string
	Servicename        string
	Applicationname    string
	Applicationversion string
	Releaseversion     string
	Componentname      string
	Securityplanid     string
	Exposedweb         string
	Experimental       string
	Userfacing         string
	Criticalinfra      string
	Sourcecontrol      string
}

// Function Name: Generate
// Function Description:
// Generate is a function that generates an EKS cluster configuration using the input parameters and a predefined template. It takes in several input parameters such as the name, instance type, owner, node groups, and tags, and generates an EKS cluster configuration using the predefined template. The generated configuration is written to the standard output.
// The parameters of the Generate function are:
// 1. name string: The name of the EKS cluster.
// 2. instancetype string: The instance type of the worker nodes in the EKS cluster.
// 3. owner string: The owner of the EKS cluster.
// 4. ngs []NodeGroup: An array of NodeGroup structs that represent the managed node groups of the EKS cluster.
// 5. tags AWSTags: A struct that represents the tags associated with the EKS cluster.
// The Generate function creates an EKSConfig struct by assigning values to its fields using the input parameters and environment variables. It then creates a template using the Eksctl constant from the templates package. Finally, the function executes the template with the EKSConfig struct as the input, and writes the generated configuration to the standard output.
// Example Usage:
// err := Generate("my-cluster", "t3.medium", "John Doe", []NodeGroup{}, AWSTags{})
//
//	if err != nil {
//	    log.Fatalf("Failed to generate EKS config: %v", err)
//	}
//
// The above example shows the usage of the Generate function to generate an EKS cluster configuration. The function takes in the name, instance type, owner, node groups, and tags as input parameters, and generates an EKS cluster configuration using the predefined template. If there is an error while generating the configuration, the function returns an error.
func Generate(name, instancetype, owner string, nodeGroups []NodeGroup, tags AWSTags) error {
	config := EKSConfig{
		ServiceArn:              os.Getenv("EKSServiceArn"),
		ClusterName:             name,
		ClusterRegion:           os.Getenv("EKSClusterRegion"),
		ClusterVersion:          os.Getenv("EKSClusterVersion"),
		ClusterInstanceType:     instancetype,
		ManagedNodeGroups:       nodeGroups,
		ClusterAMI:              os.Getenv("EKSClusterAMI"),
		InstanceRoleArn:         os.Getenv("EKSInstanceRoleArn"),
		KubeProxyVersion:        os.Getenv("EKSKubeProxyVersion"),
		CoreDNSVersion:          os.Getenv("EKSCoreDNSVersion"),
		EBSCSIVersion:           os.Getenv("EKSEBSCSIVersion"),
		PublicSubnetA:           os.Getenv("EKSPublicSubnetA"),
		PublicSubnetB:           os.Getenv("EKSPublicSubnetB"),
		PrivateSubnetA:          os.Getenv("EKSPrivateSubnetA"),
		PrivateSubnetB:          os.Getenv("EKSPrivateSubnetB"),
		SecurityGroup:           os.Getenv("EKSSecurityGroup"),
		SharedNodeSecurityGroup: os.Getenv("EKSSharedNodeSecurityGroup"),
		Tags:                    tags,
	}
	tmpl, err := template.New("eks-config").Parse(templates.Eksctl)
	if err != nil {
		return err
	}

	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, config); err != nil {
		return err
	}

	fmt.Println(rendered.String())
	return nil
}

func ProtoGenerate(model marketplace.Install_Extensions_Eks) (string, error) {
	nodeGroups := []NodeGroup{}
	tags := AWSTags{}
	config := EKSConfig{
		ServiceArn:              os.Getenv("EKSServiceArn"),
		ClusterName:             model.Clustername,
		ClusterRegion:           os.Getenv("EKSClusterRegion"),
		ClusterVersion:          os.Getenv("EKSClusterVersion"),
		ClusterInstanceType:     model.Nodegroups[0].Instancetype,
		ManagedNodeGroups:       nodeGroups,
		ClusterAMI:              os.Getenv("EKSClusterAMI"),
		InstanceRoleArn:         os.Getenv("EKSInstanceRoleArn"),
		KubeProxyVersion:        os.Getenv("EKSKubeProxyVersion"),
		CoreDNSVersion:          os.Getenv("EKSCoreDNSVersion"),
		EBSCSIVersion:           os.Getenv("EKSEBSCSIVersion"),
		PublicSubnetA:           os.Getenv("EKSPublicSubnetA"),
		PublicSubnetB:           os.Getenv("EKSPublicSubnetB"),
		PrivateSubnetA:          os.Getenv("EKSPrivateSubnetA"),
		PrivateSubnetB:          os.Getenv("EKSPrivateSubnetB"),
		SecurityGroup:           os.Getenv("EKSSecurityGroup"),
		SharedNodeSecurityGroup: os.Getenv("EKSSharedNodeSecurityGroup"),
		Tags:                    tags,
	}
	tmpl, err := template.New("eks-config").Parse(templates.Eksctl)
	if err != nil {
		return "", err
	}

	var rendered bytes.Buffer
	if err := tmpl.Execute(&rendered, config); err != nil {
		return "", err
	}

	fmt.Println(rendered.String())
	return rendered.String(), nil
}
