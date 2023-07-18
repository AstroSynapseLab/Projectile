package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/AstroSynapseLab/Projectile/clone"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login with GitHub credentials",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)
	
			fmt.Print("Enter GitHub username: ")
			username, _ := reader.ReadString('\n')
	
			fmt.Print("Enter GitHub PAT: ")
			pat, _ := reader.ReadString('\n')
	
			// Trim newline
			username = strings.TrimSpace(username)
			pat = strings.TrimSpace(pat)
	
			// Store the credentials in a struct
			creds := AuthConfig{
				GitHub: struct {
					Username string `yaml:"username"`
					Token    string `yaml:"token"`
				}{
					Username: username,
					Token:    pat,
				},
			}
	
			// Marshal the struct to YAML
			data, err := yaml.Marshal(&creds)
			if err != nil {
				fmt.Println("Failed to marshal credentials:", err)
				return
			}
	
			// Store the YAML in a file
			err = ioutil.WriteFile("./.projectile/auth.yaml", data, 0644)
			if err != nil {
				fmt.Println("Failed to store credentials:", err)
			} else {
				fmt.Println("Credentials stored successfully")
			}
		},
	}
	
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(loginCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
