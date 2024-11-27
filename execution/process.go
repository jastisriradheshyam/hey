/*
Copyright 2023 Jasti Sri Radhe Shyam

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package execution

import (
	"fmt"
	"hey/misc"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// OS Signal is monikered so that next process in the stack can be managed
	//  currently if process terminate signal is issued then following tasks will be skipped and main process stop
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	var processCloseSignalIssued = false
	go func() {
		<-osSignal
		processCloseSignalIssued = true
	}()
	for _, subTask := range task.SubTasks {
		// Skip other processes if main process got terminate signal
		if processCloseSignalIssued {
			break
		}
		if subTask.TaskType == "spawn" {
			executeCommand(subTask.SpawnInfo.Name, subTask.SpawnInfo.EnvVars, subTask.SpawnInfo.Args...)
		}
		if subTask.TaskType == "call_module" {
			modules.ParseModuleAndProcessTask(subTask.CallModuleInfo.Name)
		}
	}
}
