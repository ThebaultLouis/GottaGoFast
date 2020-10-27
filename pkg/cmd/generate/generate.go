package generate

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
)

type KitGenerator struct {
	Directory string
}

type KitConfig struct {
	Generators []*KitGenerator
}

type ProjectConfig struct {
	Kit string
}

func findProjectDirectory() (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// We look for the file ggf.json starting from the working directory
	currentDirectory := workingDirectory

	for currentDirectory != "/" {
		projectFilePath := path.Join(currentDirectory, "ggf.json")

		if _, err := os.Stat(projectFilePath); err == nil {
			// We found the file!
			return currentDirectory, nil
		}

		// We look at the parent directory
		currentDirectory = path.Dir(currentDirectory)
	}

	// The file does not exist!
	return "", errors.New("ggf.json not found")
}

func parseProjectFile(path string) (*ProjectConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := new(ProjectConfig)
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

func parseKitFile(path string) (*KitConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := new(KitConfig)
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}

func NewGenerateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Do the thing",
		Run: func(cmd *cobra.Command, args []string) {
			projectDirectory, err := findProjectDirectory()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			fmt.Println("Found ggf.json at", projectDirectory)

			projectConfig, err := parseProjectFile(path.Join(projectDirectory, "ggf.json"))
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			kitConfig, err := parseKitFile(path.Join(projectConfig.Kit, "kit.json"))
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			executor := NewExecutor(projectDirectory, projectConfig, kitConfig)

			if err = executor.Execute(); err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			fmt.Println("Done!")
		},
	}
}
