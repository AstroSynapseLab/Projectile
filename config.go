package main

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
	Project  string              `yaml:"project"`
	Services map[string][]Service `yaml:"services"`
}

type AuthConfig struct {
	GitHub struct {
		Username string `yaml:"username"`
		Token    string `yaml:"token"`
	} `yaml:"github"`
}