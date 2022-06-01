package main

import (
	"fmt"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/tagging"
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/hclparser"
)

var (
	// Used for flags.
	cfgFile       string
	path          string
	tags          []string
	subnets       []string
	secgroups     []string
	creator       string
	pocs          []string
	venue         string
	project       string
	servicearea   string
	capability    string
	component     string
	capversion    string
	release       string
	securityplan  string
	exposed       string
	experimental  string
	userfacing    string
	critinfra     string
	sourcecontrol string

	rootCmd = &cobra.Command{
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
			hclparser.Runp(path, tags, subnets, secgroups)
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

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVar(&creator, "creator", "", "The resource creator email")
	rootCmd.PersistentFlags().StringVar(&venue, "venue", "", "The venue (dev/test/prod)")
	rootCmd.PersistentFlags().StringVar(&project, "project", "", "The name of the project")
	rootCmd.PersistentFlags().StringVar(&servicearea, "servicearea", "", "The name of the Unity Service Area")
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
	rootCmd.PersistentFlags().StringSliceVarP(&tags, "tag", "t", []string{}, "A list of additional tags")
	rootCmd.PersistentFlags().StringSliceVarP(&subnets, "subnet", "s", []string{}, "A list of subnet ids")
	rootCmd.PersistentFlags().StringSliceVarP(&secgroups, "securitygroupids", "g", []string{}, "A list of security group ids")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "the path to the terraform files")
	//rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")

	//rootCmd.AddCommand(addCmd)
	//rootCmd.AddCommand(initCmd)
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
