package main

import (
	"../cmd/cobra"
)

// start cobra RootCcmd
// for more info on cobra and Rootcmd
// visit https://github.com/spf13/cobra
func main() {
	cobra.RootCmd.Execute()
}
