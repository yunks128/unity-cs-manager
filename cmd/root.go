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
	cfgFile            string
	path               string
	tags               []string
	subnets            []string
	secgroups          []string
	creator            string
	pocs               []string
	venue              string
	project            string
	servicearea        string
	capability         string
	component          string
	capversion         string
	release            string
	securityplan       string
	exposed            string
	experimental       string
	userfacing         string
	critinfra          string
	sourcecontrol      string
	eksName            string
	eksInstanceType    string
	owner              string
	managedNodeGroups  []string
	inputs             []string
	action             string
	deploymeta         string
	teardownname       string
	projectname        string
	servicename        string
	applicationname    string
	applicationversion string
	awstags            eks.AWSTags
	resourcename       string
	rootCmd            = &cobra.Command{Use: "Unity", Short: "Unity Command Line Tool", Long: ""}
	terraformcmd       = &cobra.Command{
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
			awstags = eks.AWSTags{

				Resourcename:       resourcename,
				Creatoremail:       creator,
				Pocemail:           pocs[0],
				Venue:              venue,
				Projectname:        projectname,
				Servicename:        servicename,
				Applicationname:    applicationname,
				Applicationversion: applicationversion,
				Releaseversion:     release,
				Componentname:      component,
				Securityplanid:     securityplan,
				Exposedweb:         exposed,
				Experimental:       experimental,
				Userfacing:         userfacing,
				Criticalinfra:      critinfra,
				Sourcecontrol:      sourcecontrol,
			}
			eks.Generate(eksName, eksInstanceType, owner, ngs, awstags)
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
	rootCmd.AddCommand(actionCmd)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVar(&resourcename, "resourcename", "", "The resource name")
	rootCmd.PersistentFlags().StringVar(&creator, "creator", "", "The resource creator email")
	rootCmd.PersistentFlags().StringVar(&venue, "venue", "", "The venue (dev/test/prod)")
	rootCmd.PersistentFlags().StringVar(&servicearea, "servicearea", "", "The name of the Unity Service Area")

	terraformcmd.PersistentFlags().StringSliceVarP(&tags, "tag", "t", []string{}, "A list of additional tags")
	terraformcmd.PersistentFlags().StringSliceVarP(&subnets, "subnet", "s", []string{}, "A list of subnet ids")
	terraformcmd.PersistentFlags().StringSliceVarP(&secgroups, "securitygroupids", "g", []string{}, "A list of security group ids")
	terraformcmd.PersistentFlags().StringVarP(&path, "path", "p", "", "the path to the terraform files")

	rootCmd.PersistentFlags().StringVar(&eksName, "name", "", "The EKS Cluster Name")
	eksCmd.PersistentFlags().StringVar(&eksInstanceType, "instancetype", "m5.xlarge", "The EKS Cluster Instance Type")
	rootCmd.PersistentFlags().StringVar(&owner, "owner", "u-cs", "The EKS Cluster Instance Type")
	rootCmd.PersistentFlags().StringArrayVarP(&managedNodeGroups, "managenodegroups", "", []string{}, "Managed Node Groups, comma separated,name,min,max,desired,instancetype")
	rootCmd.PersistentFlags().StringVar(&projectname, "projectname", "", "the unity project name deploying the cluster")
	rootCmd.PersistentFlags().StringVar(&servicename, "servicename", "", "the unity service name deploying the cluster")
	rootCmd.PersistentFlags().StringVar(&capability, "capability", "", "The name of the application")
	rootCmd.PersistentFlags().StringVar(&component, "component", "", "The primary type of application/runtime that will be run on this resource.")
	rootCmd.PersistentFlags().StringVar(&capversion, "capversion", "", "Version of the application.")
	rootCmd.PersistentFlags().StringVar(&release, "release", "", "Release version that the application belongs to.")
	rootCmd.PersistentFlags().StringVar(&securityplan, "securityplan", "", "The JPL security plan ID that this resource falls under.")
	rootCmd.PersistentFlags().StringVar(&exposed, "exposed", "", "Is this resource exposed to the web?")
	rootCmd.PersistentFlags().StringVar(&experimental, "experimental", "", "Is this an experimental resource? If so, it will be removed after a period of time.")

	rootCmd.PersistentFlags().StringVar(&userfacing, "userfacing", "", "Is this resource user facing? Does the user interact directly with this resource?")
	rootCmd.PersistentFlags().StringVar(&critinfra, "critinfra", "", "What is the level of criticality of the resource? This is mesaured on a scale of 5, with 5 being the most critical.")
	rootCmd.PersistentFlags().StringVar(&sourcecontrol, "sourcecontrol", "", "This should be an URL to the source code/or documentation of the software deployed on the resource.")
	rootCmd.PersistentFlags().StringSliceVar(&pocs, "pocs", []string{}, "The list of the point of contacts that is responsible for the resource is being deployed on.")

	rootCmd.MarkPersistentFlagRequired("resourcename")
	rootCmd.MarkPersistentFlagRequired("name")
	rootCmd.MarkPersistentFlagRequired("owner")
	rootCmd.MarkPersistentFlagRequired("projectname")
	rootCmd.MarkPersistentFlagRequired("servicename")
	rootCmd.MarkPersistentFlagRequired("capability")
	rootCmd.MarkPersistentFlagRequired("component")
	rootCmd.MarkPersistentFlagRequired("capversion")
	rootCmd.MarkPersistentFlagRequired("release")
	rootCmd.MarkPersistentFlagRequired("securityplan")
	rootCmd.MarkPersistentFlagRequired("exposed")
	rootCmd.MarkPersistentFlagRequired("experimental")
	rootCmd.MarkPersistentFlagRequired("userfacing")
	rootCmd.MarkPersistentFlagRequired("critinfra")
	rootCmd.MarkPersistentFlagRequired("sourcecontrol")
	rootCmd.MarkPersistentFlagRequired("pocs")
	rootCmd.MarkPersistentFlagRequired("creator")
	rootCmd.MarkPersistentFlagRequired("venue")
	rootCmd.MarkPersistentFlagRequired("servicearea")

	actionCmd.AddCommand(deployProjectCmd)
	actionCmd.AddCommand(teardownProjectCmd)
	actionCmd.AddCommand(listProjectsCmd)

	deployProjectCmd.PersistentFlags().StringVar(&deploymeta, "meta", "", "The metadata of the project to deploy")
	teardownProjectCmd.PersistentFlags().StringVar(&teardownname, "name", "", "The name of the project to teardown")

	rootCmd.AddCommand(eksCmd)
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
