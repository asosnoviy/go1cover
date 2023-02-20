package cmd

import (
	"strings"

	"github.com/asosnoviy/go1cover/internal/app"
	"github.com/asosnoviy/go1cover/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Source string
var Debugger string
var Format string
var BaseDir string

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&Source, "source", "s", "./src", "Source directory to read from")
	viper.BindPFlag("source", startCmd.Flags().Lookup("source"))
	viper.SetDefault("source", startCmd.LocalFlags().Lookup("source").DefValue)

	startCmd.Flags().StringVarP(&Debugger, "debugger", "u", "http://127.0.0.1:1550", "Debugger url")
	viper.BindPFlag("debuggerUrl", startCmd.Flags().Lookup("debugger"))
	viper.SetDefault("debuggerUrl", startCmd.LocalFlags().Lookup("debugger").DefValue)

	startCmd.Flags().StringVarP(&Format, "format", "f", "lcov, generic", "report format: lcov, generic")
	viper.BindPFlag("format", startCmd.Flags().Lookup("format"))
	viper.SetDefault("format", startCmd.LocalFlags().Lookup("format").DefValue)

	startCmd.Flags().StringVarP(&BaseDir, "baseDir", "b", "", "Project path (base dir for reports)")
	viper.BindPFlag("baseDir", startCmd.Flags().Lookup("baseDir"))

}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start coverage server",
	Long:  `This command must be used for start coverage server`,
	Run: func(cmd *cobra.Command, args []string) {

		config := &config.Config{
			Port:       viper.GetString("port"),
			SrcPath:    viper.GetString("source"),
			DebugerUrl: viper.GetString("debuggerUrl"),
			Reporters:  strings.Split(viper.GetString("format"), ","),
			RootPath:   viper.GetString("baseDir"),
		}

		context := app.Init(config)
		context.RunChan <- true
		app.Run(context)
	},
}
