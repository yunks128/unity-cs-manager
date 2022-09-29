package main

import (
	"fmt"
	//"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/actions"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/components"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/eks"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/tagging"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile           string
	path              string
	tags              []string
	subnets           []string
	secgroups         []string
	creator           string
	pocs              []string
	venue             string
	project           string
	servicearea       string
	capability        string
	component         string
	capversion        string
	release           string
	securityplan      string
	exposed           string
	experimental      string
	userfacing        string
	critinfra         string
	sourcecontrol     string
	eksName           string
	eksInstanceType   string
	owner             string
	managedNodeGroups []string
	inputs            []string
	action            string
	deploymeta        string
	teardownname      string
	projectname       string
	servicename       string

	rootCmd      = &cobra.Command{Use: "Unity", Short: "Unity Command Line Tool", Long: ""}
	terraformcmd = &cobra.Command{
		Use:   "parse",
		Short: "Parse Terraform scripts",
		Long:  `Parse Terraform scripts and add missing blocks or tags`,
		Run: func(cmd *cobra.Command, args []string) {
			validate("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)", creator, "creator")
			validate("(dev|test|prod)", venue, "venue")
			validate("([a-z0-9]+)", project, "project")
			validate("(cs|sps|ds|ads|as)", servicearea, "servicearea")
			validate("([a-z0-9]+)", capability, "capability")
			validate("^(\\d+\\.)?(\\d+\\.)?(\\*|\\d+)$", capversion, "capversion")
			validate("^G(\\d+\\.)?(\\d+\\.)?(\\*|\\d+)$", release, "release")
			validate("([a-z0-9]+)", component, "component")
			validate("([0-9]+)", securityplan, "securityplan")
			validate("([true|false]{1})", exposed, "exposed")
			validate("([true|false]{1})", experimental, "experimental")
			validate("([true|false]{1})", userfacing, "userfacing")
			validate("([0-5]{1})", critinfra, "critinfra")
			tags := tagging.GenerateMandatoryTags(creator, pocs, venue, project, servicearea, capability, component, capversion, release, securityplan, exposed, experimental, userfacing, critinfra, sourcecontrol)
			components.Runp(path, tags, subnets, secgroups)
		},
	}

	eksCmd = &cobra.Command{
		Use:   "eks",
		Short: "Generate valid EKS configs for deployment into U-CS",
		Long:  `Generate valid EKS configs using a set of input parameters to allow us to deploy easily to U-CS`,
		Run: func(cmd *cobra.Command, args []string) {
			ngs, _ := arrayToNodeGroup(managedNodeGroups)
			eks.Generate(eksName, eksInstanceType, owner, ngs, projectname, servicename)
		},
	}

	actionCmd = &cobra.Command{
		Use:   "action",
		Short: "Execute U-CS actions",
		Long:  "Execute U-CS Actions",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	deployProjectCmd = &cobra.Command{
		Use:   "deploy",
		Short: "Execute U-CS actions",
		Long:  "Execute U-CS Actions",
		Run: func(cmd *cobra.Command, args []string) {
			//actions.Execute(deploymeta)
		},
	}

	teardownProjectCmd = &cobra.Command{
		Use:   "teardown",
		Short: "Teardown U-CS actions",
		Long:  "Teardown U-CS Actions",
		Run: func(cmd *cobra.Command, args []string) {
			//actions.TearDown(teardownname)
		},
	}

	listProjectsCmd = &cobra.Command{
		Use:   "list",
		Short: "List U-CS actions",
		Long:  "List U-CS Actions",
		Run: func(cmd *cobra.Command, args []string) {
			//actions.List()
		},
	}
)

func arrayToNodeGroup(groups []string) ([]eks.NodeGroup, error) {
	ng := []eks.NodeGroup{}
	for _, g := range groups {
		s := strings.Split(g, ",")
		s1, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, err
		}
		s2, err := strconv.Atoi(s[2])
		if err != nil {
			return nil, err
		}
		s3, err := strconv.Atoi(s[3])
		if err != nil {
			return nil, err
		}
		n := eks.NodeGroup{
			s[0],
			s1,
			s2,
			s3,
			s[4],
		}
		ng = append(ng, n)
	}
	return ng, nil
}

func validate(s string, c string, name string) {
	re, err := regexp.Compile(s)
	if err != nil {
		log.Fatalf("Invalid Regex %v", s)
	}

	if !re.MatchString(c) {
		//return fmt.Errorf("invalid value: %q", flag1)
		log.Fatalf("Invalid flag %v with value %v, expected regex: %v", name, c, s)
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(terraformcmd)
	rootCmd.AddCommand(eksCmd)
	rootCmd.AddCommand(actionCmd)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	terraformcmd.PersistentFlags().StringVar(&creator, "creator", "", "The resource creator email")
	terraformcmd.PersistentFlags().StringVar(&venue, "venue", "", "The venue (dev/test/prod)")
	terraformcmd.PersistentFlags().StringVar(&project, "project", "", "The name of the project")
	terraformcmd.PersistentFlags().StringVar(&servicearea, "servicearea", "", "The name of the Unity Service Area")
	terraformcmd.PersistentFlags().StringVar(&capability, "capability", "", "The name of the application")
	terraformcmd.PersistentFlags().StringVar(&component, "component", "", "The primary type of application/runtime that will be run on this resource.")
	terraformcmd.PersistentFlags().StringVar(&capversion, "capversion", "", "Version of the application.")
	terraformcmd.PersistentFlags().StringVar(&release, "release", "", "Release version that the application belongs to.")
	terraformcmd.PersistentFlags().StringVar(&securityplan, "securityplan", "", "The JPL security plan ID that this resource falls under.")
	terraformcmd.PersistentFlags().StringVar(&exposed, "exposed", "", "Is this resource exposed to the web?")
	terraformcmd.PersistentFlags().StringVar(&experimental, "experimental", "", "Is this an experimental resource? If so, it will be removed after a period of time.")
	terraformcmd.PersistentFlags().StringVar(&userfacing, "userfacing", "", "Is this resource user facing? Does the user interact directly with this resource?")
	terraformcmd.PersistentFlags().StringVar(&critinfra, "critinfra", "", "What is the level of criticality of the resource? This is mesaured on a scale of 5, with 5 being the most critical.")
	terraformcmd.PersistentFlags().StringVar(&sourcecontrol, "sourcecontrol", "", "This should be an URL to the source code/or documentation of the software deployed on the resource.")
	terraformcmd.PersistentFlags().StringSliceVar(&pocs, "pocs", []string{}, "The list of the point of contacts that is responsible for the resource is being deployed on.")
	terraformcmd.PersistentFlags().StringSliceVarP(&tags, "tag", "t", []string{}, "A list of additional tags")
	terraformcmd.PersistentFlags().StringSliceVarP(&subnets, "subnet", "s", []string{}, "A list of subnet ids")
	terraformcmd.PersistentFlags().StringSliceVarP(&secgroups, "securitygroupids", "g", []string{}, "A list of security group ids")
	terraformcmd.PersistentFlags().StringVarP(&path, "path", "p", "", "the path to the terraform files")

	eksCmd.PersistentFlags().StringVar(&eksName, "clustername", "", "The EKS Cluster Name")
	eksCmd.PersistentFlags().StringVar(&eksInstanceType, "instancetype", "m5.xlarge", "The EKS Cluster Instance Type")
	eksCmd.PersistentFlags().StringVar(&owner, "owner", "u-cs", "The EKS Cluster Instance Type")
	eksCmd.PersistentFlags().StringArrayVarP(&managedNodeGroups, "managenodegroups", "", []string{}, "Managed Node Groups, comma separated,name,min,max,desired,instancetype")
	eksCmd.PersistentFlags().StringVar(&projectname, "projectname", "", "the unity project name deploying the cluster")
	eksCmd.PersistentFlags().StringVar(&servicename, "servicename", "", "the unity service name deploying the cluster")

	actionCmd.AddCommand(deployProjectCmd)
	actionCmd.AddCommand(teardownProjectCmd)
	actionCmd.AddCommand(listProjectsCmd)

	deployProjectCmd.PersistentFlags().StringVar(&deploymeta, "meta", "", "The metadata of the project to deploy")
	teardownProjectCmd.PersistentFlags().StringVar(&teardownname, "name", "", "The name of the project to teardown")

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
