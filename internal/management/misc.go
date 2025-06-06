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

package management

import (
	"fmt"
	configMod "hey/internal/configuration"
	utils "hey/internal/utils"
	"log"
	"os"
	"path"
)

func getUserHomeDir() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dirname, nil
}

func getConfigRootDir() (string, error) {
	homeDir, err := getUserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".hey"), nil
}

func getModulesRootDir() (string, error) {
	homeDir, err := getUserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, ".hey", "modules"), nil
}

func getConfigPath(moduleName string, configDir string) string {
	return path.Join(configDir, moduleName+".yaml")
}

func initDefaultConfig() {
	rootModulesDir, err := getModulesRootDir()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defaultModulePath := getConfigPath("default", rootModulesDir)
	data, err := configMod.GetBlankConfigYaml()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := os.WriteFile(defaultModulePath, data, 0o644); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func initModules() {
	rootModulesDir, err := getModulesRootDir()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := os.RemoveAll(rootModulesDir); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := os.Mkdir(rootModulesDir, 0o755); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	initDefaultConfig()
}
func initConfig() {
	rootConfigDir, err := getConfigRootDir()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := os.RemoveAll(rootConfigDir); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := os.Mkdir(rootConfigDir, 0o755); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	initModules()
}

func checkModuleExists(moduleName string, modulesDir string) (bool, error) {
	configPath := getConfigPath(moduleName, modulesDir)
	return utils.PathExistsByPathType(configPath, "file")
}

func GetConfigModuleBytes(moduleName string, modulesDir string) []byte {
	configPath := getConfigPath(moduleName, modulesDir)
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return configBytes
}

// Check both config dir in Home dir and default config and create both of them if not present
func CheckAndInit() {
	configDir, err := getConfigRootDir()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	modulesDir, err := getModulesRootDir()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	configCheck, err := utils.PathExistsByPathType(configDir, "dir")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	modulesCheck, err := utils.PathExistsByPathType(modulesDir, "dir")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !configCheck {
		initConfig()
	} else if !modulesCheck {
		initModules()
	} else {
		check, err := checkModuleExists("default", modulesDir)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if !check {
			initDefaultConfig()
		}
	}
}

func GetConfigModule(module string) configMod.CurrentConfigSchema {
	modulesDir, err := getModulesRootDir()
	check, err := checkModuleExists(module, modulesDir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !check {
		log.Fatal(fmt.Sprintf("%s does not exists", module))
		os.Exit(1)
	}
	configModuleBytes := GetConfigModuleBytes(module, modulesDir)
	configModuleVersion := configMod.GetConfigVersion(configModuleBytes)
	if configModuleVersion != configMod.CURRENT_CONFIG_VERSION {
		log.Fatal(fmt.Sprintf("version %d in module is not compatible to current supported version %d", configModuleVersion, configMod.CURRENT_CONFIG_VERSION))
		os.Exit(1)
	}
	return configMod.GetConfig(configModuleBytes)
}
