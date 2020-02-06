package cobra

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"path/filepath"
)

// RootCmd uses the cobra package to create a CLI for the user
// to interact with. Commands should be placed under the cobra
// package in this directory and added to RootCmd
// for more info on cobra see https://github.com/spf13/cobra
var RootCmd = &cobra.Command{
	Use: "./api_storage",
	Short: "./api_storage is used to store API keys with low level encryption locally" +
		" in the home directory of the user",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey,
		"key", "k", "",
		"Specify the key that will be used for encoding and decoding")
}

func vaultDir() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")

}
