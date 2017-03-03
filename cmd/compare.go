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
	"fmt"
	"os"
	"sync"

	"github.com/sascha-andres/go-filecomparer/app/filedb"
	"github.com/sascha-andres/go-filecomparer/app/scanner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare files to file database",
	Long: `Scans the directory and compares to data in file database.

For each change a line will be printed:

D file - file was deleted [not yet implenented]
A file - file was added
C file - file was changed

For an unchanged file, nothing will be printed`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := filedb.ConnectDB(); err != nil {
			sugar.Errorw("Error connecting to database", "err", err)
			os.Exit(4)
		}
		defer filedb.CloseDB()
		if viper.GetBool("verbose") {
			if log, err := zap.NewDevelopment(); err == nil {
				scanner.SetLogger(log)
			}
		}
		fs := make(chan (filedb.File))
		exitChannel := make(chan (bool))
		var wg sync.WaitGroup
		go scanner.Scan(fs, exitChannel, &wg)
		wg.Wait()
		for {
			select {
			case file := <-fs:
				f, err := filedb.Get(file.RelativePath)
				if err != nil {
					sugar.Errorw("Error getting file from database", "RelativePath", file.RelativePath)
				}
				if nil == f || "" == f.RelativePath {
					fmt.Printf("A %s\n", file.RelativePath)
				} else {
					if f.Hash != file.Hash {
						fmt.Printf("C %s\n", file.RelativePath)
					}
				}
			case <-exitChannel:
				return
			default:
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(compareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
