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

package misc

import "strings"

func ParseModuleTask(moduleTask string) []string {
	if len(moduleTask) == 0 {
		return []string{}
	}
	return strings.Split(moduleTask, ".")
}

// Returns module name and task name
func GetModuleAndCommandName(parsedRunArgs []string) (string, string) {
	totalArgParsedLen := len(parsedRunArgs)
	if totalArgParsedLen == 1 {
		return "default", parsedRunArgs[0]
	}

	return strings.Join(parsedRunArgs[0:(totalArgParsedLen-1)], "."), parsedRunArgs[totalArgParsedLen-1]
}

func ParseExportExcludeValues(excludeValues string) []string {
	return []string{}
}
