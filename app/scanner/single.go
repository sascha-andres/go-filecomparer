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
	"fmt"
	"os"

	"github.com/sascha-andres/go-filecomparer/app/filedb"
)

// GetFileData returns file information for a single file
func GetFileData(path string) (*filedb.File, error) {
	if ok, _ := exists(path); !ok {
		return nil, fmt.Errorf("File does not exist")
	}
	result, err := hash(path)
	if err != nil {
		return nil, err
	}
	return &filedb.File{RelativePath: path, Hash: result}, nil
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
