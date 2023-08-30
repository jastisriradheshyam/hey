package execution

import (
	config_schemas "hey/configuration/schemas"
	"hey/management"
	"runtime"
)

type SpawnInfo struct {
	Name    string
	Args    []string
	EnvVars map[string]string
}

type CallModuleInfo struct {
	Name string
}

type SubTask struct {
	TaskType       string
	SpawnInfo      *SpawnInfo
	CallModuleInfo *CallModuleInfo
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
		osName := config_schemas.EnvName(runtime.GOOS)
		if _, ok := taskConfig[osName]; !ok {
			osName = config_schemas.EnvName("default")
		}
		for _, subTaskConfig := range taskConfig[osName] {
			var subTask SubTask
			subTask.TaskType = subTaskConfig.TaskType
			if subTask.TaskType == "spawn" {
				var spawnInfo SpawnInfo
				spawnInfo.Name = subTaskConfig.SpawnInfo.Name
				spawnInfo.Args = subTaskConfig.SpawnInfo.Args
				spawnInfo.EnvVars = map[string]string{}
				for _, envVar := range subTaskConfig.SpawnInfo.EnvVars {
					key := envVar.Key
					value := envVar.Value
					spawnInfo.EnvVars[key] = value
				}
				subTask.SpawnInfo = &spawnInfo
			}
			if subTask.TaskType == "call_module" {
				var callModuleInfo CallModuleInfo
				callModuleInfo.Name = subTaskConfig.CallModuleInfo.Name
				subTask.CallModuleInfo = &callModuleInfo
			}
			subTasks = append(subTasks, &subTask)
		}

		task.SubTasks = subTasks
		(*modules)[moduleName].Tasks[taskNameConfig] = &task
	}
}
