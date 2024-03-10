package main

import (
	"github.com/jniltinho/netprivate/cmd/is-private/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
