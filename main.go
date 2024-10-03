package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"

	"github.com/danielmiessler/fabric/cli"
)

func main() {
	_, err := cli.Cli(version)
	if err != nil && !flags.WroteHelp(err) {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
