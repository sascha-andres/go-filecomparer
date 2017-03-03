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

package filedb

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // db plugin for gorm
)

var (
	// FileDatabasePath is the path to look for database
	FileDatabasePath = ".filedb"
	// DB is an instance of a database connection
	DB *gorm.DB
)

// Initialize called to create an empty database
func Initialize() error {
	sugar.Debug("Initialize")
	if nil != DB {
		DB.Close()
		DB = nil
	}
	if err := removeDatabase(); err != nil {
		return err
	}
	return ConnectDB()
}

// ConnectDB opens a database and migrates
func ConnectDB() error {
	sugar.Debug("ConnectDB")
	var err error
	DB, err = gorm.Open("sqlite3", FileDatabasePath)
	if err == nil {
		sugar.Debug("Starting migration")
		DB.AutoMigrate(&File{})
		err = DB.Error
	}
	return err
}

// CloseDB disconnects from database
func CloseDB() error {
	sugar.Debug("ConnectDB")
	return DB.Close()
}

func removeDatabase() error {
	sugar.Debug("removeDatabase")
	if ok, _ := exists(FileDatabasePath); ok {
		return os.Remove(FileDatabasePath)
	}
	return nil
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
