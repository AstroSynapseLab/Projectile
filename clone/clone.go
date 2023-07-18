package clone

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/AstroSynapseLab/Projectile/models"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/yaml.v2"
)

var (
	ErrNotProjectileRepo = errors.New("provided url is not a projectile repo")
	ErrFaildAuth         = errors.New("failed to authenticate, type `projectile login` to provide username and GitHub PAT \nvisit https://github.com/settings/tokens to generate PAT")	
)

func Do(configRepo string) error {
	fmt.Printf("cloning %s...\n\n", configRepo)
	err := cloneConfigRepo(configRepo)
	if err != nil {
		return err
	}

	config, err := readAndUnmarshalConfig()
	if err != nil {
		return err
	}

	fmt.Print("cloning services...\n\n")
	err = cloneServiceRepos(config)
	if err != nil {
		return err
	}

	fmt.Print("building environment...\n\n")
	err = copyFile("./.projectile/env/local/docker-compose.yaml", "./docker-compose.yaml")
	fmt.Print("Done!\n")
	return err
}

func cloneConfigRepo(url string) error {
	fmt.Print("loading projectile config...\n\n")
	parts := strings.Split(url, "/")
	name := parts[len(parts)-1]

	isValid := false
	if len(parts) >= 2 && parts[len(parts)-2] == "github.com" {
		isValid = true
	}

	if !isValid {
		return ErrNotProjectileRepo
	}

	configRepo := url
	if len(parts) > 4 {
		configRepo = url + "/tree/master/projectile"
	} else {
		configRepo = "https://github.com/" + name + "/projectile"
	}

	auth, err := readAuthConfig()
	if err != nil {
		return ErrFaildAuth
	}

	_, err = git.PlainClone("./.projectile", false, &git.CloneOptions{
		URL:      configRepo,
		Auth: &http.BasicAuth{
			Username: auth.GitHub.Username,
			Password: auth.GitHub.Token,
		},
	})

	// Marshal the struct to YAML
	data, err := yaml.Marshal(&auth)
	if err != nil {
		return err
	}

	// Store the YAML in a file
	err = ioutil.WriteFile("./.projectile/auth.yaml", data, 0644)
	return err
}

func readAndUnmarshalConfig() (models.Config, error) {
	data, err := ioutil.ReadFile("./.projectile/config.yaml")
	if err != nil {
		return models.Config{}, err
	}

	var config models.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return models.Config{}, err
	}

	return config, nil
}

func cloneServiceRepos(config models.Config) error {
	authConfig, err := readAuthConfig()
	if err != nil {
		fmt.Println("Failed to read auth config:", err)
		return err
	}

	for _, serviceList := range config.Services {
		for _, service := range serviceList {
			_, err := git.PlainClone(service.Path, false, &git.CloneOptions{
				URL:      service.Repo,
				Auth: &http.BasicAuth{
					Username: authConfig.GitHub.Username,
					Password: authConfig.GitHub.Token,
				},
			})
			if err != nil {
				return err
			}

			if service.Options.AddContainer {
				err = copyFile(service.Dockerfile, service.Path+"/Dockerfile")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func readAuthConfig() (models.AuthConfig, error) {
	data, err := ioutil.ReadFile("./.projectile/auth.yaml")
	if err != nil {
		return models.AuthConfig{}, err
	}

	var config models.AuthConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return models.AuthConfig{}, err
	}

	return config, nil
}

func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, input, 0644)
	return err
}
