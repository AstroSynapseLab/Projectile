package cmd

import (
	"fmt"
	"os"

	"github.com/GoLangWebSDK/Projectile/schema"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var rootCmd = &cobra.Command{
	Use:   "projectile",
	Short: "main",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	newCmd.AddCommand(newProjectCmd, newLibraryCmd, newSandboxCmd)
	rootCmd.AddCommand(newCmd)
}

func loadConfig() (schema.Config, error) {
	var config schema.Config
	var fileData []byte

	_, err := os.Stat("~/.projectile/config.yaml")
	if os.IsNotExist(err) {
		fileData, err = os.ReadFile("config.yaml")
		if err != nil {
			return config, err
		}
	} else {
		fileData, err = os.ReadFile("~/.projectile/config.yaml")
		if err != nil {
			return config, err
		}
	}

	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		return config, err
	}

	return config, err
}
