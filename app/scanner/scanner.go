// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scanner

import (
	"io/ioutil"
	"path/filepath"
	"sync"

	"os"

	"github.com/sascha-andres/go-filecomparer/app/filedb"
)

// Scan starts to walkt through all files & directories
func Scan(fileChannel chan<- filedb.File, exitChannel chan<- bool, wg *sync.WaitGroup) error {
	sugar.Debug("Scan")
	wg.Add(1)
	scanDirectory(".", fileChannel, wg)
	wg.Wait()
	exitChannel <- true
	return nil
}

func scanDirectory(directory string, fileChannel chan<- filedb.File, wg *sync.WaitGroup) {
	defer wg.Done()
	sugar.Debugw("scanDirectory", "directory", directory)
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		sugar.Errorw("Error getting files", "directory", directory)
		return
	}
	for _, file := range files {
		wg.Add(1)
		if file.Mode().IsDir() {
			scanDirectory(filepath.Join(directory, file.Name()), fileChannel, wg)
		} else {
			workOnFile(file, directory, fileChannel, wg)
		}
	}
}

func workOnFile(file os.FileInfo, directory string, fileChannel chan<- filedb.File, wg *sync.WaitGroup) {
	defer wg.Done()
	if file.Name() == filedb.FileDatabasePath {
		return
	}
	filePath := filepath.Join(directory, file.Name())
	sugar.Debugw("workOnFile", "file", filePath)
	hash, err := hash(filePath)
	if err != nil {
		sugar.Errorw("workOnFile", "file", filePath, "err", err)
	}
	fileChannel <- filedb.File{RelativePath: filePath, Hash: hash}
}
