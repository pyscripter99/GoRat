package modules

import (
	"fmt"
	"go-rat/agent/modules/command"
	helloworld "go-rat/agent/modules/hello_world"
	"go-rat/utils/types"
)

var Register Registry

type Module interface {
	Execute(task types.Task, params []types.Parameter) ([]types.Output, error)
}

type Registry struct {
	modules map[string]Module
}

func (r *Registry) RegisterModule(id string, module Module) {
	r.modules[id] = module
}

func (r *Registry) ExecuteModule(task types.Task) ([]types.Output, error) {
	module, ok := r.modules[task.ModuleID]
	if !ok {
		return []types.Output{}, fmt.Errorf("module with id: '%s' not found", task.ModuleID)
	}

	return module.Execute(task, task.Parameters)
}

func InitRegistry() {
	Register = Registry{}
	Register.modules = make(map[string]Module)

	// Register Modules
	Register.RegisterModule("hello-world", helloworld.Module{})
	Register.RegisterModule("command", command.Module{})
}
