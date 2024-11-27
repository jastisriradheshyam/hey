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

package schemas

type EnvName string

type Task map[EnvName][]SubTask

type SubTask struct {
	TaskType       string         `yaml:"type"`
	Context        interface{}    `yaml:"context"`
	SpawnInfo      SpawnInfo      `yaml:"spawn_info"`
	CallModuleInfo CallModuleInfo `yaml:"call_module"`
}

type EnvVar struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type SpawnInfo struct {
	Name    string   `yaml:"name"`
	Args    []string `yaml:"args"`
	EnvVars []EnvVar `yaml:"env_vars"`
}

type CallModuleInfo struct {
	Name string `yaml:"name"`
}

type Meta struct {
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}

type ConfigCommon struct {
	Version uint64 `yaml:"version"`
	Meta    Meta   `yaml:"meta"`
}
