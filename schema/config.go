package schema

type Options struct {
	AddContainer bool `yaml:"add_container"`
}

type Service struct {
	Name       string  `yaml:"name"`
	Path       string  `yaml:"path"`
	Repo       string  `yaml:"repo"`
	Dockerfile string  `yaml:"dockerfile"`
	Options    Options `yaml:"options"`
}

type Config struct {
	ProjectsRoot   string               `yaml:"projects_root"`
	ProjectFolders []string             `yaml:"project_folders"`
	LibsRoot       string               `yaml:"libs_root"`
	LibsFolders    []string             `yaml:"libs_folders"`
	ToolsRoot      string               `yaml:"tools_root"`
	ToolsFolders   []string             `yaml:"tools_folders"`
	WebsiteRoot    string               `yaml:"website_root"`
	WebsiteFolders []string             `yaml:"website_folders"`
	SandboxRoot    string               `yaml:"sandbox_root"`
	SandboxFolders []string             `yaml:"sandbox_folders"`
	Project        string               `yaml:"project"`
	Services       map[string][]Service `yaml:"services"`
}

type AuthConfig struct {
	GitHub struct {
		Username string `yaml:"username"`
		Token    string `yaml:"token"`
	} `yaml:"github"`
}
