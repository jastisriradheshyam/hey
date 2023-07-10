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
	rootConfigDir, err := getConfigRootDir()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defaultConfigPath := getConfigPath("default", rootConfigDir)
	data, err := configMod.GetBlankConfigYaml()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := os.WriteFile(defaultConfigPath, data, 0o644); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
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
	initDefaultConfig()
}

func checkModuleExists(moduleName string, configDir string) (bool, error) {
	configPath := getConfigPath(moduleName, configDir)
	return pathExists(configPath, "file")
}

func GetConfigModuleBytes(moduleName string, configDir string) []byte {
	configPath := getConfigPath(moduleName, configDir)
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
	check, err := pathExists(configDir, "dir")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !check {
		initConfig()
	} else {
		check, err = checkModuleExists("default", configDir)
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
	configDir, err := getConfigRootDir()
	check, err := checkModuleExists(module, configDir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if !check {
		log.Fatal(fmt.Sprintf("%s does not exists", module))
		os.Exit(1)
	}
	configModuleBytes := GetConfigModuleBytes(module, configDir)
	configModuleVersion := configMod.GetConfigVersion(configModuleBytes)
	if configModuleVersion != configMod.CURRENT_CONFIG_VERSION {
		log.Fatal(fmt.Sprintf("version %d in module is not compatible to current supported version %d", configModuleVersion, configMod.CURRENT_CONFIG_VERSION))
		os.Exit(1)
	}
	return configMod.GetConfig(configModuleBytes)
}
