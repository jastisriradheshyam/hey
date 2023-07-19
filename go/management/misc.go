package management

import (
	"fmt"
	configMod "hey/configuration"
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

// valid values for pathType are "dir" and "file"
func pathExists(path string, pathType string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		if pathType == "dir" && stat.IsDir() {
			return true, nil
		}
		if pathType == "file" {
			return true, nil
		}
		return false, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
	return pathExists(configPath, "file")
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
	configCheck, err := pathExists(configDir, "dir")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	modulesCheck, err := pathExists(modulesDir, "dir")
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
