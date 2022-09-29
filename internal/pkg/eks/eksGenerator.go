package eks

import (
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/templates"
	"os"
)
import "html/template"

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
	CoreDNSVersion          string
	SubnetConfigA           string
	SubnetConfigB           string
	SecurityGroup           string
	SharedNodeSecurityGroup string
	ClusterInstanceType     string
	ClusterOwner            string
	ManagedNodeGroups       []NodeGroup
	ServiceName             string
	ProjectName             string
}

type NodeGroup struct {
	NodeGroupName          string
	ClusterMinSize         int
	ClusterMaxSize         int
	ClusterDesiredCapacity int
	ClusterInstanceType    string
}

func Generate(name, instancetype, owner string, ngs []NodeGroup, projectname string, servicename string) error {
	sweaters := EKSConfig{
		ServiceArn:              os.Getenv("EKSServiceArn"),
		ClusterName:             name,
		ClusterRegion:           os.Getenv("EKSClusterRegion"),
		ClusterVersion:          os.Getenv("EKSClusterVersion"),
		ClusterInstanceType:     instancetype,
		ManagedNodeGroups:       ngs,
		ClusterAMI:              os.Getenv("EKSClusterAMI"),
		InstanceRoleArn:         os.Getenv("EKSInstanceRoleArn"),
		KubeProxyVersion:        os.Getenv("EKSKubeProxyVersion"),
		CoreDNSVersion:          os.Getenv("EKSCoreDNSVersion"),
		SubnetConfigA:           os.Getenv("EKSSubnetConfigA"),
		SubnetConfigB:           os.Getenv("EKSSubnetConfigB"),
		SecurityGroup:           os.Getenv("EKSSecurityGroup"),
		SharedNodeSecurityGroup: os.Getenv("EKSSharedNodeSecurityGroup"),
		ClusterOwner:            owner,
		ProjectName:             projectname,
		ServiceName:             servicename,
	}
	tmpl, err := template.New("test").Parse(templates.Eksctl)
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	return nil
}
