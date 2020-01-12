package cobra

import (
	"../../internal/storage"
	"fmt"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set a key in predefined vault",
	Run: func(cmd *cobra.Command, args []string) {
		v := storage.File(encodingKey, vaultDir())
		fmt.Println(args)
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Value set successfully.")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
