package main

import (
	"fmt"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/components"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/eks"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/tagging"
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile         string
	path            string
	tags            []string
	subnets         []string
	secgroups       []string
	creator         string
	pocs            []string
	venue           string
	project         string
	servicearea     string
	capability      string
	component       string
	capversion      string
	release         string
	securityplan    string
	exposed         string
	experimental    string
	userfacing      string
	critinfra       string
	sourcecontrol   string
	eksName         string
	eksInstanceType string
	eksMinNodes     int
	eksDesiredNodes int
	eksMaxNodes     int

	rootCmd      = &cobra.Command{Use: "Unity", Short: "Unity Command Line Tool", Long: ""}
	terraformcmd = &cobra.Command{
		Use:   "parse",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			eks.Generate(eksName, eksInstanceType, eksMinNodes, eksMaxNodes, eksDesiredNodes)
		},
	}
)

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
	eksCmd.PersistentFlags().IntVar(&eksMinNodes, "minnodes", 1, "The EKS Cluster Min Nodes")
	eksCmd.PersistentFlags().IntVar(&eksMaxNodes, "maxnodes", 3, "The EKS Cluster Max Nodes")
	eksCmd.PersistentFlags().IntVar(&eksDesiredNodes, "desirednodes", 1, "The EKS Cluster Desired Nodes")

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
