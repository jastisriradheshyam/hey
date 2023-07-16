package execution

import (
	"fmt"
	config_schemas "hey/configuration/schemas"
	"hey/management"
	"log"
	"os"
	"runtime"
)

type SpawnInfo struct {
	Name string
	Args []string
}

type SubTask struct {
	TaskType  string
	SpawnInfo *SpawnInfo
}

type Task struct {
	Name     string
	SubTasks []*SubTask
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
		var subTasks []*SubTask
		task.Name = taskNameConfig

		// Check if current exec env present in the config
		env := config_schemas.EnvName(runtime.GOOS)
		if _, ok := taskConfig[env]; !ok {
			env = config_schemas.EnvName("default_env")
		}
		for _, subTaskConfig := range taskConfig[env] {
			var subTask SubTask
			subTask.TaskType = subTaskConfig.TaskType
			if subTask.TaskType == "spawn" {
				var spawnInfo SpawnInfo
				spawnInfo.Name = subTaskConfig.SpawnInfo.Name
				spawnInfo.Args = subTaskConfig.SpawnInfo.Args
				subTask.SpawnInfo = &spawnInfo
			}
			subTasks = append(subTasks, &subTask)
		}

		task.SubTasks = subTasks
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
	for _, subTask := range task.SubTasks {
		if subTask.TaskType == "spawn" {
			executeCommand(subTask.SpawnInfo.Name, subTask.SpawnInfo.Args...)
		}
	}
}
