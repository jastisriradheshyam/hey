package schemas

type Command struct {
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

type Task struct {
	TaskType string    `yaml:"type"`
	Commands []Command `yaml:"commands"`
}

type Meta struct {
	Description string `yaml:"description"`
}

type ConfigCommon struct {
	Version uint64 `yaml:"version"`
	Meta    Meta   `yaml:"meta"`
}
