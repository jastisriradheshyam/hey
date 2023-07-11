package execution

import (
	"fmt"
	"hey/management"
	"log"
	"os"
)

type Command struct {
	Name string
	Args []string
}

type Task struct {
	Name     string
	TaskType string
	Commands *[]Command
}

type Module struct {
	Tasks map[string]*Task
}

type Modules map[string]*Module

func (module *Module) IsTaskPresent(taskName string) bool {
	_, isPresent := (*module).Tasks[taskName]
	return isPresent
}

func (modules *Modules) IsModuleLoaded(moduleName string) bool {
	_, isPresent := (*modules)[moduleName]
	return isPresent
}

func (modules *Modules) LoadModule(moduleName string) {
	configModule := management.GetConfigModule(moduleName)
	var module Module
	(*modules)[moduleName] = &module
	(*modules)[moduleName].Tasks = make(map[string]*Task)
	for taskNameConfig, taskConfig := range configModule.Tasks {
		var task Task
		task.Name = taskNameConfig
		task.TaskType = taskConfig.TaskType
		if taskConfig.TaskType == "command" {
			var commands []Command
			for _, command := range taskConfig.Commands {
				commands = append(commands, Command{Name: command.Name, Args: command.Args})
			}
			task.Commands = &commands
		}
		(*modules)[moduleName].Tasks[taskNameConfig] = &task
	}
}

func (modules *Modules) ProcessTask(moduleName string, taskName string) {
	if isModuleLoaded := modules.IsModuleLoaded(moduleName); !isModuleLoaded {
		modules.LoadModule(moduleName)
	}
	if isTaskPresent := (*modules)[moduleName].IsTaskPresent(taskName); !isTaskPresent {
		log.Fatal(fmt.Sprintf("Task: %s is not present in the module: %s", taskName, moduleName))
		os.Exit(1)
	}
	task := *(*modules)[moduleName].Tasks[taskName]
	executeCommand((*task.Commands)[0].Name, (*task.Commands)[0].Args...)
}
