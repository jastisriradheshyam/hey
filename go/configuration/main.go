package configuration

import (
	schemas "hey/configuration/schemas"

	yaml "gopkg.in/yaml.v3"
)

const CURRENT_CONFIG_VERSION = uint64(1)

type CurrentConfigSchema schemas.ConfigV1

func GetConfigVersion(config []byte) uint64 {
	var configCommon schemas.ConfigCommon
	err := yaml.Unmarshal(config, &configCommon)
	if err != nil {
		panic(err)
	}
	return configCommon.Version
}

func GetConfig(config []byte) CurrentConfigSchema {
	var configModule CurrentConfigSchema
	err := yaml.Unmarshal(config, &configModule)
	if err != nil {
		panic(err)
	}
	return configModule
}

func GetBlankConfigYaml() ([]byte, error) {
	var config CurrentConfigSchema
	config.Version = CURRENT_CONFIG_VERSION
	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		return []byte{}, err
	}
	return yamlBytes, nil
}
