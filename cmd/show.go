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

DB updated hash
CURRENT current hash

Exit code is 0 for identical hashes, else 1`,
	Run: func(cmd *cobra.Command, args []string) {
		activateLogger()
		if err := filedb.ConnectDB(); err != nil {
			sugar.Errorw("Error connecting to database", "err", err)
			os.Exit(4)
		}
		defer filedb.CloseDB()
		if "" == viper.GetString("show.file") {
			fmt.Println("Please provide the file")
			os.Exit(4)
		}
		dbFile, err := filedb.Get(viper.GetString("show.file"))
		if err != nil {
			sugar.Errorw("Error getting file information from database", "err", err)
			os.Exit(4)
		}
		fmt.Printf("DB %v %s\n", dbFile.UpdatedAt, dbFile.Hash)
		f, err := scanner.GetFileData(viper.GetString("show.file"))
		if err != nil {
			sugar.Errorw("Error getting file information", "err", err)
			os.Exit(4)
		}
		fmt.Printf("CURRENT current %s\n", f.Hash)
		if f.Hash == dbFile.Hash {
			os.Exit(0)
		}
		os.Exit(1)
	},
}

func init() {
	RootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("file", "f", "", "File to commit")
	viper.BindPFlag("show.file", showCmd.Flags().Lookup("file"))

}
