package cobra

import (
	"../../internal/logger"
	"../../internal/storage"
	"fmt"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del [API]",
	Short: "Deletes the specific API:key pair from the vault",
	Long: "Deletes the API:key pair from the vault. " +
		" API to be deleted is specified in args. If fails" +
		" corresponding message is given.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger.GeneralLogger.Println("Delete command called")
		v := storage.File(encodingKey, vaultDir())
		apiName := args[0]
		fmt.Println("Attempting to delete pair")
		v.DeletePair(apiName)
	},
}

func init() {
	RootCmd.AddCommand(delCmd)
}
