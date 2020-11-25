package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/daithi-coombes/go-api-eth-skeleton/pkg/bot"
)

var cfgFile string
var conv string
var protocol string
var org string
var endpoint string

var rootCmd = &cobra.Command{
	Use:   "go-api-eth-skeleton",
	Short: "A bot to pull data from DAO smart contracts",
	Long:  `TECommons DAO Chat Bot`,
	Run: func(cmd *cobra.Command, args []string) {

		// start bot
		// TODO: use channels to decouple dao instance from bot. bot.Run() should just run server/handlers
		bot.Run(cmd)

	},
}

// Execute The main function for executing the cli package (Cobra).
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
	rootCmd.PersistentFlags().StringVar(&conv, "conviction", "", "The conviction voting address")
	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "https://rpc.xdaichain.com", "The url of the node to connect to")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".go-api-eth-skeleton")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
