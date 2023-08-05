package utils

import "os"

// valid values for pathType are "dir" and "file"
func PathExistsByPathType(path string, pathType string) (bool, error) {
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

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
