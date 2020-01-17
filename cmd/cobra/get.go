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
		apiName := args[0]
		logger.GeneralLogger.Printf("Retrieved key for %s", apiName)
		value, err := v.Get(apiName)
		if err != nil {
			fmt.Println("API not found, could not retrieve key")
			logger.ErrorLogger.Fatalf("could not get key, failed with error"+
				"%s", err)
		}
		fmt.Printf("%s \n retrieved for %s", value, apiName)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
