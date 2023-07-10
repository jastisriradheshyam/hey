package schemas

type ConfigV1 struct {
	Version uint64          `yaml:"version"`
	Meta    Meta            `yaml:"meta"`
	Tasks   map[string]Task `yaml:"tasks"`
}
