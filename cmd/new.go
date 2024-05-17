package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project, library or sandbox",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New command called")
	},
}

var newProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create new project, requires project name as argument",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New project called")

		config, err := loadConfig()
		if err != nil {
			fmt.Println("Error loading config", err)
			return
		}

		projectName := args[0]
		fullPath := fmt.Sprintf("%s%s", loadFullPath(config.ProjectsRoot), projectName)

		os.MkdirAll(fullPath, 0755)

		//create subfolders
		for _, folder := range config.ProjectFolders {
			os.MkdirAll(fmt.Sprintf("%s%s%s", fullPath, "/", folder), 0755)
		}

		fmt.Println("project created at: ", fullPath)
	},
}

var newLibraryCmd = &cobra.Command{
	Use:   "lib",
	Short: "Create new library",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New library called")

		config, err := loadConfig()
		if err != nil {
			fmt.Println("Error loading config", err)
			return
		}

		libraryName := args[0]
		fullPath := fmt.Sprintf("%s%s", loadFullPath(config.LibsRoot), libraryName)

		os.MkdirAll(fullPath, 0755)

		//create subfolders
		for _, folder := range config.LibsFolders {
			os.MkdirAll(fmt.Sprintf("%s%s%s", fullPath, "/", folder), 0755)
		}

		fmt.Println("library created at: ", fullPath)
	},
}

var newSandboxCmd = &cobra.Command{
	Use:   "sandbox",
	Short: "Create new sandbox",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New sandbox called")

		config, err := loadConfig()
		if err != nil {
			fmt.Println("Error loading config", err)
			return
		}

		sandboxName := args[0]
		fullPath := fmt.Sprintf("%s%s", loadFullPath(config.SandboxRoot), sandboxName)

		os.MkdirAll(fullPath, 0755)

		//create subfolders
		for _, folder := range config.SandboxFolders {
			os.MkdirAll(fmt.Sprintf("%s%s%s", fullPath, "/", folder), 0755)
		}
	},
}

var newToolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Create new tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("New tool called")

		config, err := loadConfig()

		if err != nil {
			fmt.Println("Error loading config", err)
			return
		}

		toolName := args[0]
		fullPath := fmt.Sprintf("%s%s", loadFullPath(config.ToolsRoot), toolName)

		os.MkdirAll(fullPath, 0755)

		//create subfolders
		for _, folder := range config.ToolsFolders {
			os.MkdirAll(fmt.Sprintf("%s%s%s", fullPath, "/", folder), 0755)
		}

		fmt.Println("tool created at: ", fullPath)
	},
}

func loadFullPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Unable to find the user's home directory:", err)
			return path
		}
		return strings.Replace(path, "~", home, 1)
	}
	return path
}
