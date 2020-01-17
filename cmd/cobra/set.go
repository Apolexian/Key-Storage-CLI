package cobra

import (
	"../../internal/logger"
	"../../internal/storage"
	"fmt"
	"github.com/spf13/cobra"
)

// setCmd uses the cobra package to create a set command
// this sets a value and its associated key
// example usage: ./script set example-api example-key
// will set the example-key for the example-api in the
// homedir/.secrets file
var setCmd = &cobra.Command{
	Use:   "set [value] [key] where value is the API name and key is API key",
	Short: "create/set an entry in the secrets file for given API",
	Long: "set takes two parameters, value and key. Value is the " +
		"name of the API (or other value) that is to be stored" +
		"and is used for lookup. Key is the value to be stored." +
		"If value already exists, key will be overridden. The pair" +
		"is stored in homedir/.secrets.",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		logger.GeneralLogger.Println("Set called by user")
		v := storage.File(encodingKey, vaultDir())
		apiName, key := args[0], args[1]
		logger.GeneralLogger.Printf("User set %s for %s", key, apiName)
		err := v.Set(apiName, key)
		if err != nil {
			logger.ErrorLogger.Fatalf("could not get key, failed with error"+
				"%s", err)
		}
		fmt.Printf("Key set successfully under name: %s.", apiName)
		logger.GeneralLogger.Println("Set successfully")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
