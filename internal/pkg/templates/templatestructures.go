package templates

const Eksctl = `
apiVersion: eksctl.io/v1alpha5
kind: ClusterConfig

iam:
  serviceRoleARN: {{ .ServiceArn }}
  withOIDC: false

metadata:
  name: {{ .ClusterName }}
  region: {{ .ClusterRegion }}
  version: "{{ .ClusterVersion }}"
  tags:
    service: "{{ .ServiceName }}"
    project: "{{ .ProjectName }}"
    owner:   "{{ .ClusterOwner }}"

addons:
  - name: kube-proxy
    version: {{ .KubeProxyVersion }}
    tags:
      service: "{{ .ServiceName }}"
      project: "{{ .ProjectName }}"
      owner: "{{ .ClusterOwner }}"
  - name: coredns
    version: {{ .CoreDNSVersion }}
  - name: aws-ebs-csi-driver
    version: {{ .EBSCSIVersion}}
vpc:
  subnets:
    private:
      {{ .SubnetConfigA }}
      {{ .SubnetConfigB }}
  securityGroup: {{ .SecurityGroup }}
  sharedNodeSecurityGroup: {{ .SharedNodeSecurityGroup }}
  manageSharedNodeSecurityGroupRules: false

managedNodeGroups:
{{- range $key, $value := .ManagedNodeGroups }}
  - name: {{ $value.NodeGroupName }}NodeGroup
    minSize: {{ $value.ClusterMinSize }}
    maxSize: {{ $value.ClusterMaxSize }}
    desiredCapacity: {{ $value.ClusterDesiredCapacity }}
    instanceType: {{ $value.ClusterInstanceType }}
    ami: {{ $.ClusterAMI }}
    tags:
      service: "{{ $.ServiceName }}"
      project: "{{ $.ProjectName }}"
      owner:  "{{ $.ClusterOwner }}"
    iam:
      instanceRoleARN: {{ $.InstanceRoleArn }}
    privateNetworking: true
    overrideBootstrapCommand: |
      #!/bin/bash
      /etc/eks/bootstrap.sh {{ $.ClusterName }}
{{- end }}
`
