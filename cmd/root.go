package cmd

import "github.com/spf13/cobra"

var (
	dbPath  string
	rootCmd = &cobra.Command{
		Use:   "leveldb-cli",
		Short: "A CLI tool for interacting with LeveDB",
		Long: `A command line interface tool for performing operations on LevelDB databases.
Complete documentation is available at https://github.com/coldinke/leveldb-cli`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbPath, "path", "", "Path to LevelDB database (required)")
	rootCmd.MarkPersistentFlagRequired("path")
}
