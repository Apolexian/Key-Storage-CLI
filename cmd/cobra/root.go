package cobra

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"path/filepath"
)

var RootCmd = &cobra.Command{
	Use:   "vault",
	Short: "vault is used to store api keys with low level encryption locally",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey,
		"key", "k", "",
		"key that will be used for encoding and decoding")
}

func vaultDir() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")

}
