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

	"github.com/sascha-andres/go-filecomparer/app/filedb"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if a file is unchanged",
	Long: `Use this command to check if a file is unchanged

Return code 0 if unchanged
Return code 1 if changed
Return code 2 if deleted
Return code 3 if added
Return code 4 on error`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := filedb.ConnectDB(); err != nil {
			sugar.Errorw("Error connecting to database", "err", err)
			os.Exit(4)
		}
		defer filedb.CloseDB()
	},
}

func init() {
	RootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
