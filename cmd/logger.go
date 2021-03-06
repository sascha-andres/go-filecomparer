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

package cmd

import (
	"log"

	"github.com/sascha-andres/go-filecomparer/app/filedb"
	"github.com/sascha-andres/go-filecomparer/app/scanner"
	"github.com/spf13/viper"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal("Error creating logger")
	}
	sugar = logger.Sugar()
}

// SetLogger is used to set another logger
func SetLogger(newLogger *zap.Logger) {
	logger = newLogger
	sugar = newLogger.Sugar()
}

func activateLogger() {
	if viper.GetBool("verbose") {
		if log, err := zap.NewDevelopment(); err == nil {
			scanner.SetLogger(log)
			filedb.SetLogger(log)
			SetLogger(log)
		}
	}
}
