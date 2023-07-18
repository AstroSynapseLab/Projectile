package clone

import (
	"io/ioutil"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
)

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

func Do(configRepo string) error {
	// Clone the config repo
	_, err := git.PlainClone("./.projectile", false, &git.CloneOptions{
		URL:      configRepo,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	// Read the config file
	data, err := ioutil.ReadFile("./.projectile/config.yaml")
	if err != nil {
		return err
	}

	// Unmarshal the YAML
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	// Clone the service repos
	for _, serviceList := range config.Services {
		for _, service := range serviceList {
			_, err := git.PlainClone(service.Path, false, &git.CloneOptions{
				URL:      service.Repo,
				Progress: os.Stdout,
			})
			if err != nil {
				return err
			}

			// Copy the Dockerfile if the add_container option is true
			if service.Options.AddContainer {
				err = copyFile(service.Dockerfile, service.Path+"/Dockerfile")
				if err != nil {
					return err
				}
			}
		}
	}

	err = copyFile("./.projectile/env/local/docker-compose.yaml", "./docker-compose.yaml")
	if err != nil {
		return err
	}

	return nil
}

func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, input, 0644)
	return err
}
