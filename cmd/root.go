package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/daithi-coombes/go-api-eth-skeleton/pkg/dao"
)

var cfgFile string
var protocol string
var org string
var endpoint string

var rootCmd = &cobra.Command{
	Use:   "go-api-eth-skeleton",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		dao, err := dao.NewDAO("aragon", common.HexToAddress(org), endpoint)
		if err != nil {
			log.Fatalf("Error creating DAO instance: %s\n", err)
		}

		total, err := dao.TotalProposals()
		if err != nil {
			log.Fatalf("Error getting total proposals: %s\n", err)
		}

		fmt.Printf("Total Proposals: %d\n", total)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-api-eth-skeleton.yaml)")
	rootCmd.PersistentFlags().StringVar(&protocol, "protocol", "aragon", "the protocol to use (default is Aragon)")
	rootCmd.PersistentFlags().StringVar(&org, "organization", "", "The organisation address")
	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "https://rpc.xdaichain.com", "The url of the node to connect to")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-api-eth-skeleton" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-api-eth-skeleton")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
