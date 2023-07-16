package schemas

type EnvName string

type Task map[EnvName][]SubTask

type SubTask struct {
	TaskType  string    `yaml:"type"`
	SpawnInfo SpawnInfo `yaml:"spawn_info"`
}

type SpawnInfo struct {
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

type Meta struct {
	Description string `yaml:"description"`
}

type ConfigCommon struct {
	Version uint64 `yaml:"version"`
	Meta    Meta   `yaml:"meta"`
}
