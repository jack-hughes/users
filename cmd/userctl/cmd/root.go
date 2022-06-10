package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	host string
	port string
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "userctl",
		Short: "A friendly CLI for interacting with the users service",
	}

	rootCmd.PersistentFlags().StringVarP(&host, "host", "a", "localhost", "host of the gRPC Server")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "5355", "port of the gRPC Server")
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		log.Fatalf("failed to bind flag values: %v", err)
	}

	rootCmd.AddCommand(createCmd())
	rootCmd.AddCommand(updateCmd())
	rootCmd.AddCommand(deleteCmd())
	rootCmd.AddCommand(listCmd())

	return rootCmd
}
