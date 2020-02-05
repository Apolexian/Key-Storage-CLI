package cobra

import (
	"fmt"

	"../../internal/logger"
	"../../internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all currently stored keys in vault",
	Long: "provides a list of all currently stored" +
		" API:key pairs stored inside the vault, if" +
		" vault is not found error is shown",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		logger.GeneralLogger.Println("User called list command")
		v := storage.File(encodingKey, vaultDir())
		fmt.Println("All keys currently stored in vault: ")
		err := v.GetAllPairs()
		if err != nil {
			fmt.Println("Something went wrong when listing pairs...")
			logger.ErrorLogger.Fatalf("Error on retrieving list pairs")
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
