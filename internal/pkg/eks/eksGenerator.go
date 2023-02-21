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
	EBSCSIVersion           string
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
	Tags                    AWSTags
}

type NodeGroup struct {
	NodeGroupName          string
	ClusterMinSize         int
	ClusterMaxSize         int
	ClusterDesiredCapacity int
	ClusterInstanceType    string
}

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

func Generate(name, instancetype, owner string, ngs []NodeGroup, tags AWSTags) error {
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
		EBSCSIVersion:           os.Getenv("EKSEBSCSIVersion"),
		SubnetConfigA:           os.Getenv("EKSSubnetConfigA"),
		SubnetConfigB:           os.Getenv("EKSSubnetConfigB"),
		SecurityGroup:           os.Getenv("EKSSecurityGroup"),
		SharedNodeSecurityGroup: os.Getenv("EKSSharedNodeSecurityGroup"),
		Tags:                    tags,
	}
	tmpl, err := template.New("test").Parse(templates.Eksctl)
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	return nil
}
