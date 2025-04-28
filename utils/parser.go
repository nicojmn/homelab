/*
This file parse the main compose file and create a new compose file for each service in the directory services/
*/

package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	yaml "gopkg.in/yaml.v3"
)

type ComposeFile struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image         string   `yaml:"image"`
	ContainerName string   `yaml:"container_name"`
	Commands      []string `yaml:"command,omitempty"`
	Environment   any      `yaml:"environment,omitempty"`
	Ports         []string `yaml:"ports,omitempty"`
	Labels        []string `yaml:"labels,omitempty"`
	Volumes       []string `yaml:"volumes,omitempty"`
	Networks      []string `yaml:"networks,omitempty"`
	Restart       string   `yaml:"restart,omitempty"`
	CapAdd        []string `yaml:"cap_add,omitempty"`
	User          string   `yaml:"user,omitempty"`
	DependsOn     []string `yaml:"depends_on,omitempty"`
	Expose        []string `yaml:"expose,omitempty"`
	Pid           string   `yaml:"pid,omitempty"`
	Deploy        any      `yaml:"deploy,omitempty"`
}

func sanitizeFileName(name string) string {
	re := regexp.MustCompile(`[^\w\-.]`)
	return re.ReplaceAllString(name, "_")
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	composePath := "../docker-compose.yml"
	servicesDir := "../services"

	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	logger.Info().Msg("Starting the parser")

	_, err := os.Stat(servicesDir)
	if os.IsNotExist(err) {
		logger.Info().Msg("services directory not found")
		err = os.MkdirAll(servicesDir, os.ModePerm)
		if err != nil {
			logger.Fatal().Err(err).Msg("Error creating services directory")
		}
		logger.Info().Msg("services directory created successfully")
	}

	_, err = os.Stat(composePath)
	if os.IsNotExist(err) {
		logger.Fatal().Msg("docker-compose.yml file not found")
	}

	file, err := os.ReadFile(composePath)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error reading docker-compose.yml file")
	}

	compose := ComposeFile{}

	err = yaml.Unmarshal([]byte(file), &compose)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error parsing docker-compose.yml")
	}
	logger.Info().Msg("Parsed docker-compose.yml file successfully")

	for _, service := range compose.Services {

		composeFile := ComposeFile{
			Services: map[string]Service{
				service.ContainerName: service,
			},
		}

		data, err := yaml.Marshal(&composeFile)

		if err != nil {
			logger.Fatal().Err(err).Msg("Error marshalling service to yaml")
		}

		filePath := filepath.Join(servicesDir, sanitizeFileName("compose-"+strings.ToLower(service.ContainerName))+".yml")

		file, _ := os.Create(filePath)
		defer file.Close()

		_, err = file.Write(data)
		if err != nil {
			logger.Fatal().Err(err).Msg("Error writing service to file")
		}
		logger.Info().Msg("Service " + service.ContainerName + " created successfully")
	}
}
