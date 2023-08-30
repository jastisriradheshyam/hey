package execution

import (
	"fmt"
	"hey/misc"
	"log"
	"os"
)

func (modules *Modules) ParseModuleAndProcessTask(moduleTask string) {
	parsedModuleTask := misc.ParseModuleTask(moduleTask)
	if len(parsedModuleTask) <= 0 {
		log.Fatal("Run Arguments are not present")
		os.Exit(1)
	}
	module, taskName := misc.GetModuleAndCommandName(parsedModuleTask)
	modules.ProcessTask(module, taskName)
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
			executeCommand(subTask.SpawnInfo.Name, subTask.SpawnInfo.EnvVars, subTask.SpawnInfo.Args...)
		}
		if subTask.TaskType == "call_module" {
			modules.ParseModuleAndProcessTask(subTask.CallModuleInfo.Name)
		}
	}
}
