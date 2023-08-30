package management

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"hey/utils"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const defaultExportFileName = "config.tar.gz"

func Export(exportPath string, excludeModules []string) error {
	var err error
	exportAbsolutePath, err := determineExportPath(exportPath)
	if err != nil {
		return err
	}
	modulesRootDir, err := getModulesRootDir()
	if err != nil {
		return err
	}
	filesContext, err := os.ReadDir(modulesRootDir)
	if err != nil {
		return err
	}
	var files []string
	var excludeModulesMap = make(map[string]bool)
	for _, v := range excludeModules {
		excludeModulesMap[v] = true
	}

	for _, file := range filesContext {
		fileNameList := strings.Split(file.Name(), ".")
		length := len(fileNameList)
		fileBaseWithoutExt := length - 1
		if fileBaseWithoutExt == 0 { // there will be at least blank string at index 0
			fileBaseWithoutExt = 1
		}
		fileNameWithoutExt := strings.Join(fileNameList[0:fileBaseWithoutExt], "")
		if !excludeModulesMap[fileNameWithoutExt] {
			files = append(files, path.Join(modulesRootDir, file.Name()))
		}
	}

	// Create output file
	out, err := os.Create(exportAbsolutePath)
	if err != nil {
		log.Fatalln("Error writing archive:", err)
	}
	defer out.Close()

	// Create the archive and write the output to the "out" Writer
	err = createArchive(files, out)
	if err != nil {
		log.Fatalln("Error creating archive:", err)
	}

	fmt.Println("Archive created successfully : " + exportAbsolutePath)
	return nil
}

func determineExportPath(exportPath string) (string, error) {
	var err error
	if exportPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return path.Join(cwd, defaultExportFileName), nil
	}
	if filepath.IsAbs(exportPath) {
		return getExportPath(exportPath)
	}
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	derivedPath := path.Join(cwd, exportPath)
	return getExportPath(derivedPath)
}

func getExportPath(exportPath string) (string, error) {
	pathExists, err := utils.PathExists(exportPath)
	if err != nil {
		return "", err
	}

	if pathExists {
		stat, err := os.Stat(exportPath)
		if err != nil {
			return "", err
		}
		if stat.IsDir() {
			return path.Join(exportPath, defaultExportFileName), nil
		}
		return "", errors.New("file already present on the provided path: " + exportPath)
	}
	dir, _ := filepath.Split(exportPath)
	if dir == "" {
		return exportPath, nil
	}
	parentPathExists, err := utils.PathExists(dir)
	if err != nil {
		return "", err
	}
	if !parentPathExists {
		return "", errors.New("parent directory of the file path does not exists. Parent Path: " + dir)
	}
	return exportPath, nil
}

func createArchive(files []string, buf io.Writer) error {
	// Create new Writers for gzip and tar
	// These writers are chained. Writing to the tar writer will
	// write to the gzip writer which in turn will write to
	// the "buf" writer
	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Iterate over files and add them to the tar archive
	for _, file := range files {
		err := addToArchive(tw, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}
