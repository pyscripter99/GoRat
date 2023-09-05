package helloworld

import (
	"fmt"
	"go-rat/utils/output"
	"go-rat/utils/types"
)

type Module struct{}

func (m Module) Execute(task types.Task, params []types.Parameter) ([]types.Output, error) {
	name, err := types.GetParameter("name", params)
	if err != nil {
		return []types.Output{}, err
	}
	fmt.Printf("Hello, %s!\n", name.Value)
	// Log output
	output.LogInfo(task.ID, "")
	return []types.Output{}, nil
}
