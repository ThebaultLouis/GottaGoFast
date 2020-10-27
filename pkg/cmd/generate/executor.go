package generate

import (
	"github.com/otiai10/copy"
	"os"
	"path"
)

type Executor struct {
	projectDirectory string
	projectConfig    *ProjectConfig
	kitConfig        *KitConfig

	destinationDirectory string
}

func NewExecutor(projectDirectory string, projectConfig *ProjectConfig, kitConfig *KitConfig) *Executor {
	destinationDirectory := path.Join(projectDirectory, "generated")

	return &Executor{
		projectDirectory: projectDirectory,
		projectConfig:    projectConfig,
		kitConfig:        kitConfig,

		destinationDirectory: destinationDirectory,
	}
}

func (executor *Executor) Execute() error {
	if err := os.RemoveAll(executor.destinationDirectory); err != nil {
		return err
	}

	for _, generator := range executor.kitConfig.Generators {
		if err := executor.runSimpleGenerator(generator); err != nil {
			return err
		}
	}

	return nil
}

func (executor *Executor) runSimpleGenerator(generator *KitGenerator) error {
	sourceDirectory := path.Join(executor.projectDirectory, executor.projectConfig.Kit, generator.Directory)

	return copy.Copy(sourceDirectory, executor.destinationDirectory)
}
