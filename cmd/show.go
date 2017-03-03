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

	"github.com/sascha-andres/go-filecomparer/app/filedb"
	"github.com/sascha-andres/go-filecomparer/app/scanner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details about a file",
	Long: `Print out details of a file such as stored hash and current hash.

Output looks like this:

INDB hash updated
FILE hash

Exit code is   0 for identical hashes
Exit code is   1 if file is not in database and in filesystem
Exit code is   2 if file is not in database and not in filesystem
Exit code is   3 if file is not in filesystem but in database
Exit code is 100 on any other errors`,
	Run: func(cmd *cobra.Command, args []string) {
		activateLogger()
		if "" == viper.GetString("show.file") {
			fmt.Println("Please provide the file")
			os.Exit(100)
		}
		if err := filedb.ConnectDB(); err != nil {
			sugar.Errorw("Error connecting to database", "err", err)
			os.Exit(100)
		}
		defer filedb.CloseDB()
		f, dbFile, err := getData()
		if err != nil {
			sugar.Errorw("Error connecting to database", "err", err)
			os.Exit(100)
		}
		returnCodeHandling(f, dbFile)
		os.Exit(0)
	},
}

func getData() (*filedb.File, *filedb.File, error) {
	dbFile, err := filedb.Get(viper.GetString("show.file"))
	if err != nil {
		sugar.Errorw("Error getting file information from database", "err", err)
		os.Exit(100)
	}
	if dbFile != nil && dbFile.ID != 0 {
		fmt.Printf("INDB %s %v\n", dbFile.Hash, dbFile.UpdatedAt)
	}
	f, err := scanner.GetFileData(viper.GetString("show.file"))
	if err != nil {
		sugar.Errorw("Error getting file information", "err", err)
		os.Exit(100)
	}
	return f, dbFile, nil
}
func returnCodeHandling(f *filedb.File, dbFile *filedb.File) {
	if f != nil {
		fmt.Printf("FILE %s\n", f.Hash)
		if f.Hash == dbFile.Hash {
			os.Exit(0)
		}
	}
	if dbFile == nil || dbFile.ID == 0 {
		if f == nil {
			os.Exit(2)
		} else {
			os.Exit(1)
		}
	} else {
		if f == nil {
			os.Exit(3)
		}
	}
}

func init() {
	RootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("file", "f", "", "File to commit")
	viper.BindPFlag("show.file", showCmd.Flags().Lookup("file"))
}
