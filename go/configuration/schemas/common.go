package schemas

type EnvName string

type Task map[EnvName][]SubTask

type SubTask struct {
	TaskType  string      `yaml:"type"`
	Context   interface{} `yaml:"context"`
	SpawnInfo SpawnInfo   `yaml:"spawn_info"`
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

type Meta struct {
	Description string `yaml:"description"`
}

type ConfigCommon struct {
	Version uint64 `yaml:"version"`
	Meta    Meta   `yaml:"meta"`
}
