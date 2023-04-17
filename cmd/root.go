package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	//"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/actions"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/components"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/eks"
	"github.com/unity-sds/unity-cs-terraform-transformer/internal/pkg/tagging"

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
	// inputs             []string
	// action             string
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
			_ = eks.Generate(eksName, eksInstanceType, owner, ngs, awstags)
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
			// actions.Execute(deploymeta)
		},
	}

	teardownProjectCmd = &cobra.Command{
		Use:   "teardown",
		Short: "Teardown U-CS actions",
		Long:  "Teardown U-CS Actions",
		Run: func(cmd *cobra.Command, args []string) {
			// actions.TearDown(teardownname)
		},
	}

	listProjectsCmd = &cobra.Command{
		Use:   "list",
		Short: "List U-CS actions",
		Long:  "List U-CS Actions",
		Run: func(cmd *cobra.Command, args []string) {
			// actions.List()
		},
	}
)

// Function Name: arrayToNodeGroup
// Function Signature: func arrayToNodeGroup(groups []string) ([]eks.NodeGroup, error)
// Function Description:
// arrayToNodeGroup is a function that converts an array of strings representing EKS NodeGroups to an array of NodeGroup structs. It takes an input parameter 'groups' which is an array of strings where each string represents a NodeGroup in the format "NodeGroupName,ClusterMinSize,ClusterMaxSize,ClusterDesiredCapacity,ClusterInstanceType". The function parses each string in the array to extract the individual fields and creates a NodeGroup struct with those values. It then appends the newly created NodeGroup to the resulting slice of NodeGroups. The final output of the function is the resulting slice of NodeGroups.
//
// Function Parameters:
// 1. groups []string : An array of strings representing EKS NodeGroups. Each string is in the format "NodeGroupName,ClusterMinSize,ClusterMaxSize,ClusterDesiredCapacity,ClusterInstanceType".
//
// Return Values:
// 1. []eks.NodeGroup : A slice of NodeGroup structs that represents the NodeGroups in the input array of strings.
// 2. error : An error that is returned if any of the input strings are not in the correct format or if there is an error while parsing the values.

// Function Logic:
// The function starts by initializing an empty slice of NodeGroups called 'ng'.
// It then iterates through each string in the input array 'groups' using a for loop and extracts the individual fields from the string using the 'strings.Split()' function.
// The resulting slice of strings is then parsed using the 'strconv.Atoi()' function to convert the string values to integers where necessary.
// The function then creates a new NodeGroup struct using the parsed values and appends it to the 'ng' slice.
// Finally, the function returns the resulting slice of NodeGroups 'ng' and any error that occurred during the processing of the input array.
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
			NodeGroupName:          s[0],
			ClusterMinSize:         s1,
			ClusterMaxSize:         s2,
			ClusterDesiredCapacity: s3,
			ClusterInstanceType:    s[4],
		}
		ng = append(ng, n)
	}
	return ng, nil
}

// Function Signature: func validate(s string, c string, name string)
// Function Description:
// validate is a function that takes in three parameters - a regular expression string 's', a string 'c', and a name string 'name'. The function compiles the regular expression string 's' using the 'regexp.Compile()' function and checks if the string 'c' matches the compiled regular expression. If the string 'c' does not match the regular expression, the function logs a fatal error with details about the invalid flag and its value.
//
// Function Parameters:
// 1. s string: A regular expression string that is used to validate the string 'c'.
// 2. c string: A string value that needs to be validated against the regular expression 's'.
// 3. name string: A string that represents the name of the flag being validated.
//
// Return Values: This function does not return any values, it logs a fatal error if the string 'c' does not match the regular expression 's'.
//
// Function Logic:
// The function starts by compiling the regular expression string 's' using the 'regexp.Compile()' function. If there is an error while compiling the regular expression, the function logs a fatal error with the details of the invalid regular expression.
// Next, the function uses the 're.MatchString()' function to check if the string 'c' matches the compiled regular expression. If the string 'c' does not match the regular expression, the function logs a fatal error with details about the invalid flag and its value.
//
// Example Usage:
//
// s := `^[a-zA-Z]+$`
// name := "flag1"
// value := "invalid123"
//
// validate(s, value, name)
//
// Output:
//
// Fatal error: Invalid flag flag1 with value invalid123, expected regex: ^[a-zA-Z]+$
//
// The above example shows the usage of the validate function to validate a string 'value' against a regular expression string 's'. In this case, the regular expression expects only alphabets and does not allow any numbers or special characters. Since the string 'value' contains numbers, the function logs a fatal error with details about the invalid flag and its value.
func validate(s string, c string, name string) {
	re, err := regexp.Compile(s)
	if err != nil {
		log.Fatalf("Invalid Regex %v", s)
	}

	if !re.MatchString(c) {
		// return fmt.Errorf("invalid value: %q", flag1)
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
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
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
	rootCmd.PersistentFlags().StringVar(&applicationname, "applicationname", "", "Application name")
	rootCmd.PersistentFlags().StringVar(&applicationversion, "applicationversion", "", "Application version")

	_ = rootCmd.MarkPersistentFlagRequired("resourcename")
	_ = rootCmd.MarkPersistentFlagRequired("name")
	_ = rootCmd.MarkPersistentFlagRequired("owner")
	_ = rootCmd.MarkPersistentFlagRequired("projectname")
	_ = rootCmd.MarkPersistentFlagRequired("servicename")
	_ = rootCmd.MarkPersistentFlagRequired("capability")
	_ = rootCmd.MarkPersistentFlagRequired("component")
	_ = rootCmd.MarkPersistentFlagRequired("capversion")
	_ = rootCmd.MarkPersistentFlagRequired("release")
	_ = rootCmd.MarkPersistentFlagRequired("securityplan")
	_ = rootCmd.MarkPersistentFlagRequired("exposed")
	_ = rootCmd.MarkPersistentFlagRequired("experimental")
	_ = rootCmd.MarkPersistentFlagRequired("userfacing")
	_ = rootCmd.MarkPersistentFlagRequired("critinfra")
	_ = rootCmd.MarkPersistentFlagRequired("sourcecontrol")
	_ = rootCmd.MarkPersistentFlagRequired("pocs")
	_ = rootCmd.MarkPersistentFlagRequired("creator")
	_ = rootCmd.MarkPersistentFlagRequired("venue")
	_ = rootCmd.MarkPersistentFlagRequired("servicearea")
	_ = rootCmd.MarkPersistentFlagRequired("applicationname")
	_ = rootCmd.MarkPersistentFlagRequired("applicationversion")
	actionCmd.AddCommand(deployProjectCmd)
	actionCmd.AddCommand(teardownProjectCmd)
	actionCmd.AddCommand(listProjectsCmd)

	deployProjectCmd.PersistentFlags().StringVar(&deploymeta, "meta", "", "The metadata of the project to deploy")
	teardownProjectCmd.PersistentFlags().StringVar(&teardownname, "name", "", "The name of the project to teardown")

	rootCmd.AddCommand(eksCmd)
}

// Function Name: initConfig
// Function Signature: func initConfig()
//
// Function Description:
// initConfig is a function that initializes the configuration for the application. It checks if the configuration file path is specified using a command-line flag 'cfgFile'. If the flag is set, the function sets the configuration file path to the value of the flag using the 'viper.SetConfigFile()' function.
// If the flag is not set, the function attempts to find the home directory of the user using the 'os.UserHomeDir()' function. It then sets the configuration path to a file called ".cobra.yaml" in the user's home directory using the 'viper.AddConfigPath()', 'viper.SetConfigType()', and 'viper.SetConfigName()' functions.
// The function then reads the configuration values from the file using the 'viper.ReadInConfig()' function. If the configuration file is found and read successfully, the function prints a message indicating the configuration file that is being used.
//
// Function Parameters: This function does not take any input parameters.
//
// Return Values: This function does not return any values.
//
// Function Logic:
// The function starts by checking if the 'cfgFile' flag is set. If the flag is set, the function sets the configuration file path to the value of the flag using the 'viper.SetConfigFile()' function.
// If the flag is not set, the function finds the home directory of the user using the 'os.UserHomeDir()' function. It then sets the configuration path to a file called ".cobra.yaml" in the user's home directory using the 'viper.AddConfigPath()', 'viper.SetConfigType()', and 'viper.SetConfigName()' functions.
// The function then reads the configuration values from the file using the 'viper.ReadInConfig()' function. If the configuration file is found and read successfully, the function prints a message indicating the configuration file that is being used.
// Finally, the function enables automatic environment variable binding using the 'viper.AutomaticEnv()' function. This allows the application to read configuration values from environment variables as well.
//
// Example Usage:
// The initConfig function is typically called at the beginning of an application to initialize the configuration values. It is usually called from the main function of the application.
// func main() {
//     // Initialize configuration
//     initConfig()

//     // Other application logic
//     ...
// }
//
// The above example shows the usage of the initConfig function in a typical application. The function is called at the beginning of the main function to initialize the configuration values. The application logic can then use the configuration values by reading them from the 'viper' package.
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
