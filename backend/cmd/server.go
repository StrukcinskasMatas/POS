package cmd

import (
	"os"

	"pos/internal/handler"
	"github.com/spf13/cobra"
)

var (
	port string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start POS web server",
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleHTTP(port)
		os.Exit(0)
	},
}

func newServerCmd() *cobra.Command {
	cmd := serverCmd
	cmd.Flags().StringVarP(&port, "port", "p", "8080", "server port")

	return cmd
}

func init() {
	rootCmd.AddCommand(newServerCmd())
}
