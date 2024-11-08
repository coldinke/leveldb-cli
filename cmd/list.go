package cmd

import (
	"fmt"

	"github.com/coldinke/leveldb-cli/pkg/db"
	"github.com/spf13/cobra"
)

var (
	prefix  string
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all key-value pairs",
		RunE: func(cmd *cobra.Command, args []string) error {
			manager, err := db.NewManager(dbPath)
			if err != nil {
				return err
			}
			defer manager.Close()

			results, err := manager.List(prefix)
			if err != nil {
				return err
			}

			if len(results) == 0 {
				if prefix != "" {
					fmt.Printf("No keys found with prefix '%s'\n", prefix)
				} else {
					fmt.Println("Database is empty")
				}
				return nil
			}

			for _, kv := range results {
				fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
			}
			fmt.Printf("\nTotal: %d records\n", len(results))
			return nil
		},
	}
)

func init() {
	listCmd.Flags().StringVar(&prefix, "prefix", "", "List keys with specific prefix")
	rootCmd.AddCommand(listCmd)
}
