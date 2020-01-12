package cobra

import (
	"../../internal/storage"
	"fmt"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a key from predefined vault",
	Run: func(cmd *cobra.Command, args []string) {
		v := storage.File(encodingKey, vaultDir())
		fmt.Println(args)
		key := args[0]
		value, err := v.Get(key)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s=%s", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
