package main

import (
	"fmt"
	"os"

	"github.com/AstroSynapseLab/Projectile/clone"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "projectile",
	Short: "main",
}

func init() {
	
	// creates .projectile directory
	// inits option -m --mono will create a monolith app setup
	// defualt creates services cluster
	// init requires a github project or a workspace
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initilaze new project",
		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}

	// clones projectile folder
	// pulls all the repos from config
	// copies all env and config files
	cloneCmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone existing project",
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			err := clone.Do(url)
			if err != nil {
				fmt.Println(err)
			}
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
