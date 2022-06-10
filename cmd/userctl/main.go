package main

import (
	"os"

	"github.com/jack-hughes/users/cmd/userctl/cmd"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
