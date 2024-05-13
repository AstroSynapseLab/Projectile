package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New command called")
	},
}
