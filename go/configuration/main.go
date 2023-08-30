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
	for taskKey := range configModule.Tasks {
		for envName := range configModule.Tasks[taskKey] {
			for index := range configModule.Tasks[taskKey][envName] {
				contextBytes, err := yaml.Marshal(configModule.Tasks[taskKey][envName][index].Context)
				if err != nil {
					panic(err)
				}
				// If uncommented this will make type based context (e.g. spawn_info) usage in config,
				// else overwrite the type based context with context value, which will make consistent
				// way to set configuration
				// if len(contextBytes) == 0 {
				// 	continue
				// }
				if configModule.Tasks[taskKey][envName][index].TaskType == "spawn" {
					var spawnInfo schemas.SpawnInfo
					err = yaml.Unmarshal(contextBytes, &spawnInfo)
					if err != nil {
						panic(err)
					}
					configModule.Tasks[taskKey][envName][index].SpawnInfo = spawnInfo
				}
				if configModule.Tasks[taskKey][envName][index].TaskType == "call_module" {
					var callModuleInfo schemas.CallModuleInfo
					err = yaml.Unmarshal(contextBytes, &callModuleInfo)
					if err != nil {
						panic(err)
					}
					configModule.Tasks[taskKey][envName][index].CallModuleInfo = callModuleInfo
				}
				// Remove context data to clean up the memory
				var blankInterface interface{}
				configModule.Tasks[taskKey][envName][index].Context = blankInterface
			}
		}
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
