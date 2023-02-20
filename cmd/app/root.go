package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// This is used for config file
	cfgFile string
	port    string
	rootCmd = &cobra.Command{
		Use:     "go1cover",
		Short:   "go1cover: coverage collector for 1C:Enterprise",
		Long:    `go1cover: line coverage collector for 1C:Enterprice through http debug server`,
		Version: `1.0.1`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".go1cover.yaml", "config file name")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "4040", "coverage server port")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", rootCmd.PersistentFlags().Lookup("port").DefValue)

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {

		viper.SetConfigType("yaml")
		viper.SetConfigName(".go1cover")
	}

	viper.SetEnvPrefix("COVER")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
