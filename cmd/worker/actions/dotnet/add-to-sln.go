package dotnet

import (
	"context"
	"fmt"
	"github.com/PiotrFerenc/mash2/cmd/worker/actions"
	"github.com/PiotrFerenc/mash2/internal/Container"
	"github.com/PiotrFerenc/mash2/internal/types"
)

func AddProjectToSolution() actions.Action {
	return &addToSln{
		projectName: actions.Property{
			Name:        "projectName",
			Type:        actions.Text,
			Description: "The path to the project that needs to be added to the solution",
			DisplayName: "Project Path",
			Validation:  "required",
		},
		projectPath: actions.Property{
			Name:        "projectPath",
			Type:        actions.Text,
			Description: "The path to the project that needs to be added to the solution",
			DisplayName: "Project Path",
			Validation:  "required",
		},
		solutionPath: actions.Property{
			Name:        "solutionPath",
			Type:        actions.Text,
			Description: "The path to the solution where the project needs to be added",
			DisplayName: "Solution Path",
			Validation:  "required",
		},
		containerId: actions.Property{
			Name:        "ContainerId",
			Type:        "string",
			Description: "Unique identifier of the container.",
			DisplayName: "Container ID",
			Validation:  "Required",
		},
	}
}

type addToSln struct {
	projectPath  actions.Property
	solutionPath actions.Property
	projectName  actions.Property
	containerId  actions.Property
}

func (a addToSln) Inputs() []actions.Property {
	return []actions.Property{
		a.solutionPath, a.projectPath, a.projectName,
	}
}

func (a addToSln) Outputs() []actions.Property {
	return make([]actions.Property, 0)
}

func (a addToSln) GetCategoryName() string {
	return "dotnet"
}

func (a addToSln) Execute(process types.Process) (types.Process, error) {
	ctx := context.Background()
	projectPath, err := a.projectPath.GetStringFrom(&process)
	if err != nil {
		return process, err
	}
	solutionPath, err := a.solutionPath.GetStringFrom(&process)
	if err != nil {
		return process, err
	}
	projectName, err := a.projectName.GetStringFrom(&process)
	if err != nil {
		return process, err
	}

	envSolutionPath := fmt.Sprintf("SOLUTION_PATH=%s", solutionPath)
	envProjectPath := fmt.Sprintf("PROJECT_PATH=%s", projectPath)
	imageName := "dotnet-solutionFile"
	vol := fmt.Sprintf("/dashboard/appuser/%s:/data", projectName)
	containerId, err := Container.StartContainer(imageName, []string{envSolutionPath, envProjectPath}, []string{vol}, ctx)

	process.SetString(a.containerId.Name, containerId)
	return process, nil
}
