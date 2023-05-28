package stringer

import (
	"fmt"

	"github.com/k2font/cli-application-test/package/stringer"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args[0]
		res, kind := stringer.Inspect(i, false)
		fmt.Printf("'%s' has a %d %s.\n", i, res, kind)
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
