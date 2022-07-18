package templates

const Eksctl = `
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig

iam:
  serviceRoleARN: {{ .ServiceArn }}

metadata:
  name: {{ .ClusterName }}
  region: {{ .ClusterRegion }}
  version: "{{ .ClusterVersion }}"
  tags:
    {{ .ClusterOwner }}



managedNodeGroups:
  - name: {{ .ClusterName }}NodeGroup
    minSize: {{ .ClusterMinSize }}
    maxSize: {{ .ClusterMaxSize }}
    desiredCapacity: {{ .ClusterDesiredCapacity }}
	instanceType: {{ .ClusterInstanceType }}
    ami: {{ .ClusterAMI }}
    iam:
      instanceRoleARN: {{ .InstanceRoleArn }}
    privateNetworking: true
    overrideBootstrapCommand: |
      #!/bin/bash
      /etc/eks/bootstrap.sh {{ .ClusterName }}
addons:
  - name: kube-proxy
    version: {{ .KubeProxyVersion }}
  - name: coredns
    version: {{ .CoreDNSVersion }}


vpc:
  subnets:
    private:
      {{ .SubnetConfigA }}
      {{ .SubnetConfigB }}
  securityGroup: {{ .SecurityGroup }}
  sharedNodeSecurityGroup: {{ .SharedNodeSecurityGroup }}
  manageSharedNodeSecurityGroupRules: false
`
