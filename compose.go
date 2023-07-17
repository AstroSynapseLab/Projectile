package docker

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Service struct {
	ContainerName string `yaml:"container_name"`
	Repo          string `yaml:"repo"`
	Dockerfile    string `yaml:"dockerfile"`
	Ports         []string `yaml:"ports"`
}

type Config struct {
	Project  string              `yaml:"project"`
	Services map[string][]Service `yaml:"services"`
	Database struct {
		Image          string   `yaml:"image"`
		ContainerName  string   `yaml:"container_name"`
		Environment    []string `yaml:"environment"`
		Ports          []string `yaml:"ports"`
	} `yaml:"database"`
	Network struct {
		Name string `yaml:"asai-network"`
	} `yaml:"network"`
}

type DockerCompose struct {
	Version  string `yaml:"version"`
	Services map[string]interface{} `yaml:"services"`
	Volumes  map[string]interface{} `yaml:"volumes"`
	Networks map[string]interface{} `yaml:"networks"`
}

func buildDockerCompose() {
	// Read config file
	file, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the config file
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err)
	}

	// Create DockerCompose structure
	dockerCompose := DockerCompose{
		Version:  "3",
		Services: make(map[string]interface{}),
		Volumes:  make(map[string]interface{}),
		Networks: make(map[string]interface{}),
	}

	for serviceName, serviceDetails := range config.Services {
		for _, serviceDetail := range serviceDetails {
			dockerCompose.Services[serviceDetail.ContainerName] = map[string]interface{}{
				"container_name": serviceDetail.ContainerName,
				"build": map[string]string{
					"context":    fmt.Sprintf("./%s-%s", config.Project, serviceName),
					"dockerfile": serviceDetail.Dockerfile,
				},
				"ports":    serviceDetail.Ports,
				"volumes":  []string{fmt.Sprintf("./%s-%s:/app", config.Project, serviceName)},
				"depends_on": []string{config.Database.ContainerName},
				"networks":  []string{config.Network.Name},
			}
		}
	}

	if config.Database.ContainerName != "" {
		dbVolume := fmt.Sprintf("%s-db:/var/lib/postgresql/data", config.Project)
		dockerCompose.Services[config.Database.ContainerName] = map[string]interface{}{
			"image":          config.Database.Image,
			"container_name": config.Database.ContainerName,
			"environment":    config.Database.Environment,
			"volumes":        []string{dbVolume},
			"ports":          config.Database.Ports,
			"restart":        "unless-stopped",
			"networks":       []string{config.Network.Name},
		}
		dockerCompose.Volumes[dbVolume] = nil
	}

	dockerCompose.Networks[config.Network.Name] = nil

	// Marshal the DockerCompose
	dockerComposeData, err := yaml.Marshal(&dockerCompose)
	if err != nil {
		fmt.Println(err)
	}

	// Write the DockerCompose data to a file
	err = ioutil.WriteFile("./docker-compose.yaml", dockerComposeData, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("docker-compose.yaml file has been created.")
}
