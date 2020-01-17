package cobra

import (
	"../../internal/logger"
	"../../internal/storage"
	"fmt"
	"github.com/spf13/cobra"
)

// getCmd uses the cobra package to create a get command
// this returns the key associated with the value that
// is input, example usage: ./script get example-api
// will return the key associated with example-api
var getCmd = &cobra.Command{
	Use: "get [value] where value is the name of the API for which" +
		"the associated key needs to be retrieved",
	Short: "retrieve the API key from the vault",
	Long: "get takes the API name that was loaded into" +
		"the vault using set and retrieves the key associated" +
		"with it. If the key does not exist an error is shown",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger.GeneralLogger.Println("Get called by user")
		v := storage.File(encodingKey, vaultDir())
		key := args[0]
		logger.GeneralLogger.Printf("Retrieved key for %s", key)
		value, err := v.Get(key)
		if err != nil {
			logger.ErrorLogger.Println("could not get key")
			panic(err)
		}
		fmt.Printf("%s = %s", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
