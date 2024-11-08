package cmd

import (
	"fmt"

	"github.com/coldinke/leveldb-cli/pkg/db"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [key]",
	Short: "Delete key-value pair",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := db.NewManager(dbPath)
		if err != nil {
			return err
		}
		defer manager.Close()

		if err := manager.Delete(args[0]); err != nil {
			return err
		}

		fmt.Printf("Successfully deleted key '%s'\n", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
