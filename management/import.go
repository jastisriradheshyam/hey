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
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

// Import from a path
func Import(importPath string) error {
	if importPath == "" {
		return errors.New("No import path defined")
	}

	f, err := os.Open(importPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tarReader := tar.NewReader(gzf)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("extracting the tar file failed: %s", err.Error())
		}

		modulesRootDir, err := getModulesRootDir()
		if err != nil {
			return err
		}
		CheckAndInit()
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(path.Join(modulesRootDir, header.Name), 0755); err != nil {
				log.Fatalf("creating directory failed: %s", err.Error())
			}
		case tar.TypeReg:
			// TODO: validate and ask if module already exists and add it
			outFile, err := os.Create(path.Join(modulesRootDir, header.Name))
			if err != nil {
				log.Fatalf("creating file failed: %s", err.Error())
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("copying file content from tar failed: %s", err.Error())
			}
		// skip symlink not supported by hey
		// case tar.TypeSymlink:
		default:
			log.Fatalf(
				"while extracting there was unknown file type: %c in %s",
				header.Typeflag,
				header.Name)
		}
	}
	return nil
}
