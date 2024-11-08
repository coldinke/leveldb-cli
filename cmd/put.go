package cmd

import (
	"fmt"

	"github.com/coldinke/leveldb-cli/pkg/db"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put [key] [value]",
	Short: "Put key-value pair",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := db.NewManager(dbPath)
		if err != nil {
			return err
		}
		defer manager.Close()

		if err := manager.Put(args[0], args[1]); err != nil {
			return err
		}

		fmt.Printf("Successfully put key '%s' with value '%s'\n", args[0], args[1])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
