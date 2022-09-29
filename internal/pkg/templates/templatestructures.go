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
    service: "{{ .ServiceName }}"
    project: "{{ .ProjectName }}"

managedNodeGroups:
{{ range $value := .NodeGroups }}
  - name: {{ .value.NodeGroupName }}NodeGroup
    minSize: {{ .value.ClusterMinSize }}
    maxSize: {{ .value.ClusterMaxSize }}
    desiredCapacity: {{ .value.ClusterDesiredCapacity }}
    instanceType: {{ .value.ClusterInstanceType }}
    ami: {{ .ClusterAMI }}
    tags:
	  service: "{{ .ServiceName }}"
	  project: "{{ .ProjectName }}"
    iam:
      instanceRoleARN: {{ .InstanceRoleArn }}
    privateNetworking: true
    overrideBootstrapCommand: |
      #!/bin/bash
      /etc/eks/bootstrap.sh {{ .ClusterName }}
{{ end }}
addons:
  - name: kube-proxy
    version: {{ .KubeProxyVersion }}
    tags:
      service: "{{ .ServiceName }}"
      project: "{{ .ProjectName }}"
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
