package cmd

import (
	"fmt"

	"github.com/coldinke/leveldb-cli/pkg/db"
	"github.com/spf13/cobra"
)

var existsCmd = &cobra.Command{
	Use:   "exists [key]",
	Short: "Check if key exists",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := db.NewManager(dbPath)
		if err != nil {
			return err
		}
		defer manager.Close()

		exists, err := manager.Exists(args[0])
		if err != nil {
			return err
		}

		if exists {
			fmt.Printf("Key '%s' exists\n", args[0])
		} else {
			fmt.Printf("Key '%s' does not exist\n", args[0])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(existsCmd)
}
