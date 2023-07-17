package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "projectile",
	Short: "main",
}

func init() {
	
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initilaze new project",
		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}

	cloneCmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone existing project",
		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(cloneCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
