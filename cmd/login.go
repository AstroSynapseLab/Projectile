package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AstroSynapseLab/Projectile/schema"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with GitHub credentials",
	Run: func(cmd *cobra.Command, args []string) {
		_ = os.MkdirAll("./.projectile", 0755)

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("enter GitHub username: ")
		username, _ := reader.ReadString('\n')

		fmt.Print("enter GitHub PAT: ")
		pat, _ := reader.ReadString('\n')

		// Trim newline
		username = strings.TrimSpace(username)
		pat = strings.TrimSpace(pat)

		// Store the credentials in a struct
		auth := schema.AuthConfig{
			GitHub: struct {
				Username string `yaml:"username"`
				Token    string `yaml:"token"`
			}{
				Username: username,
				Token:    pat,
			},
		}

		// Marshal the struct to YAML
		data, err := yaml.Marshal(&auth)
		if err != nil {
			fmt.Println("login failed:", err)
			return
		}

		// Store the YAML in a file
		err = os.WriteFile("./.projectile/auth.yaml", data, 0644)
		if err != nil {
			fmt.Println("login failed:", err)
		} else {
			fmt.Println("login successfull")
		}
	},
}
