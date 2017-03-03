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
	"os"

	"fmt"

	"github.com/sascha-andres/go-filecomparer/app/filedb"
	"github.com/sascha-andres/go-filecomparer/app/scanner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Accept a changed file",
	Long:  `Takes a changed file and updates the file database`,
	Run: func(cmd *cobra.Command, args []string) {
		activateLogger()
		if err := filedb.ConnectDB(); err != nil {
			sugar.Errorw("Error connecting to database", "err", err)
			os.Exit(4)
		}
		defer filedb.CloseDB()
		if "" == viper.GetString("commit.file") {
			fmt.Println("Please provide the file")
			os.Exit(4)
		}
		f, err := scanner.GetFileData(viper.GetString("commit.file"))
		if err != nil {
			sugar.Errorw("Error getting file information", "err", err)
			os.Exit(4)
		}
		err = f.Save()
		if err != nil {
			sugar.Errorw("Error saving file information", "err", err)
			os.Exit(4)
		}
	},
}

func init() {
	RootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringP("file", "f", "", "File to commit")
	viper.BindPFlag("commit.file", commitCmd.Flags().Lookup("file"))
}
