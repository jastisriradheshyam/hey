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
