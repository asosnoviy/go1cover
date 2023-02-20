package cmd

import (
	coverserver "github.com/asosnoviy/go1cover/internal/coverServer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop coverage server",
	Long:  `This command must be used for stop coverage server and generate report`,
	Run: func(cmd *cobra.Command, args []string) {
		coverserver.NewClient(viper.GetString("port")).Stop()

	},
}
