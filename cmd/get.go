package cmd

import (
	"fmt"

	"github.com/coldinke/leveldb-cli/pkg/db"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get value by key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		manager, err := db.NewManager(dbPath)
		if err != nil {
			return err
		}
		defer manager.Close()

		value, err := manager.Get(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Value: %s\n", value)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
