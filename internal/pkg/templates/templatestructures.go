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
    unity-name: "{{ .Tags.Resourcename }}"
    unity-creator: "{{ .Tags.Creatoremail }}"
    unity-poc: "{{ .Tags.Pocemail }}"
    unity-venue: "{{ .Tags.Venue }}"
    unity-project: "{{ .Tags.Projectname }}"
    unity-service-area: "{{ .Tags.Servicename }}"
    unity-capability: "{{ .Tags.Applicationname }}"
    unity-capversion: "{{ .Tags.Applicationversion }}"
    unity-release: "{{ .Tags.Releaseversion }}"
    unity-component: "{{ .Tags.Componentname }}"
    unity-security-plan-id: "{{ .Tags.Securityplanid }}"
    unity-exposed-web: "{{ .Tags.Exposedweb }}"
    unity-experimental: "{{ .Tags.Experimental }}"
    unity-user-facing: "{{ .Tags.Userfacing }}"
    unity-crit-infra: "{{ .Tags.Criticalinfra }}"
    unity-source-control: "{{ .Tags.Sourcecontrol }}"

addons:
  - name: kube-proxy
    version: {{ .KubeProxyVersion }}
    tags:
      unity-name: "{{ .Tags.Resourcename }}"
      unity-creator: "{{ .Tags.Creatoremail }}"
      unity-poc: "{{ .Tags.Pocemail }}"
      unity-venue: "{{ .Tags.Venue }}"
      unity-project: "{{ .Tags.Projectname }}"
      unity-service-area: "{{ .Tags.Servicename }}"
      unity-capability: "{{ .Tags.Applicationname }}"
      unity-capversion: "{{ .Tags.Applicationversion }}"
      unity-release: "{{ .Tags.Releaseversion }}"
      unity-component: "{{ .Tags.Componentname }}"
      unity-security-plan-id: "{{ .Tags.Securityplanid }}"
      unity-exposed-web: "{{ .Tags.Exposedweb }}"
      unity-experimental: "{{ .Tags.Experimental }}"
      unity-user-facing: "{{ .Tags.Userfacing }}"
      unity-crit-infra: "{{ .Tags.Criticalinfra }}"
      unity-source-control: "{{ .Tags.Sourcecontrol }}"
  - name: coredns
    version: {{ .CoreDNSVersion }}
  - name: aws-ebs-csi-driver
    version: {{ .EBSCSIVersion}}
vpc:
  subnets:
{{if .PrivateSubnetA}}
    private:
      {{ .PrivateSubnetA }}
      {{ .PrivateSubnetB }}
{{end}}
{{if .PublicSubnetA}}
    public:
      {{ .PublicSubnetA }}
      {{ .PublicSubnetB }}
{{end}}
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
    amiFamily: AmazonLinux2
    tags:
      unity-name: "{{ $.Tags.Resourcename }}"
      unity-creator: "{{ $.Tags.Creatoremail }}"
      unity-poc: "{{ $.Tags.Pocemail }}"
      unity-venue: "{{ $.Tags.Venue }}"
      unity-project: "{{ $.Tags.Projectname }}"
      unity-service-area: "{{ $.Tags.Servicename }}"
      unity-capability: "{{ $.Tags.Applicationname }}"
      unity-capversion: "{{ $.Tags.Applicationversion }}"
      unity-release: "{{ $.Tags.Releaseversion }}"
      unity-component: "{{ $.Tags.Componentname }}"
      unity-security-plan-id: "{{ $.Tags.Securityplanid }}"
      unity-exposed-web: "{{ $.Tags.Exposedweb }}"
      unity-experimental: "{{ $.Tags.Experimental }}"
      unity-user-facing: "{{ $.Tags.Userfacing }}"
      unity-crit-infra: "{{ $.Tags.Criticalinfra }}"
      unity-source-control: "{{ $.Tags.Sourcecontrol }}"
    iam:
      instanceRoleARN: {{ $.InstanceRoleArn }}
    privateNetworking: false
    overrideBootstrapCommand: |
      #!/bin/bash
      /etc/eks/bootstrap.sh {{ $.ClusterName }}
{{- end }}
`
