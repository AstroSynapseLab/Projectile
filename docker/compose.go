package docker

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type ServiceDetails struct {
	ContainerName string `yaml:"container_name"`
	Repo          string `yaml:"repo"`
	Dockerfile    string `yaml:"dockerfile"`
	Ports         []string `yaml:"ports"`
}

type Config struct {
	Project  string                         `yaml:"project"`
	Services map[string][]map[string]string `yaml:"services"`
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
	Version  string                    `yaml:"version"`
	Services map[string]interface{}    `yaml:"services"`
	Volumes  map[string]interface{}    `yaml:"volumes"`
	Networks map[string]struct{}       `yaml:"networks"`
}

func Build() {
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
		Networks: make(map[string]struct{}),
	}

	for serviceName, serviceDetails := range config.Services {
		service := make(map[string]interface{})
		for _, detail := range serviceDetails {
			if containerName, ok := detail["container_name"]; ok {
				service["container_name"] = containerName
			}
			if dockerfile, ok := detail["dockerfile"]; ok {
				service["build"] = map[string]string{
					"context":    fmt.Sprintf("./%s-%s", config.Project, serviceName),
					"dockerfile": dockerfile,
				}
			}
			if ports, ok := detail["ports"]; ok {
				service["ports"] = []string{ports}
			}
			service["volumes"] = []string{fmt.Sprintf("./%s-%s:/app", config.Project, serviceName)}
			service["depends_on"] = []string{config.Database.ContainerName}
			service["networks"] = []string{config.Network.Name}
		}
		dockerCompose.Services[serviceName] = service
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
		dockerCompose.Volumes["db-data"] = struct{}{}
	}

	dockerCompose.Networks[config.Network.Name] = struct{}{}

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
