package cobra

import (
	"../../internal/logger"
	"../../internal/storage"
	"fmt"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear [API]",
	Short: "Deletes everything from vault",
	Long: "Clears all API:key pairs from the" +
		"vault.",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		logger.GeneralLogger.Println("Clear command called")
		v := storage.File(encodingKey, vaultDir())
		fmt.Println("Attempting to clear")
		err := v.DeleteAll()
		if err != nil {
			fmt.Println("Failed to clear")
			logger.ErrorLogger.Fatalf("Failed to clear: %s", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(clearCmd)
}
