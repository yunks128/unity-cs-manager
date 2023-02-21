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
    unity-name: "{{ .Tags.ResourceName }}"
    unity-creator: "{{ .Tags.CreatorEmail }}"
    unity-poc: "{{ .Tags.POCEmail }}"
    unity-venue: "{{ .Tags.Venue }}"
    unity-project: "{{ .Tags.ProjectName }}"
    unity-service-area: "{{ .Tags.ServiceName }}"
    unity-capability: "{{ .Tags.ApplicationName }}"
    unity-capversion: "{{ .Tags.ApplicationVersion }}"
    unity-release: "{{ .Tags.ReleaseVersion }}"
    unity-component: "{{ .Tags.ComponentName }}"
    unity-security-plan-id: "{{ .Tags.SecurityPlanID }}"
    unity-exposed-web: "{{ .Tags.ExposedWeb }}"
    unity-experimental: "{{ .Tags.Experimental }}"
    unity-user-facing: "{{ .Tags.UserFacing }}"
    unity-crit-infra: "{{ .Tags.CriticalInfra }}"
    unity-source-control: "{{ .Tags.SourceControl }}"

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
    amiFamily: AmazonLinux2
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
