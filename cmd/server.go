package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var httpServerCommand = &cobra.Command{
	RunE:  runHTTPServer,
	Use:   "http_server",
	Short: "to run http server",
}

func runHTTPServer(_ *cobra.Command, _ []string) error {
	fmt.Println("==runHTTPServer()")
	return nil
}
