package stringer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "stringer",
	Version: version,
	Short:   "stringer - a simple CLI to transform and inspect strings",
	Long:    "このCLIアプリケーションは、ShoichiroがCobraの使い方を学ぶために書いた習作です",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "CLIアプリケーション実行中にエラーが発生しました", err)
		os.Exit(1)
	}
}
