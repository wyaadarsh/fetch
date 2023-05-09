/*
CopyRight (c) 2023
License: MIT
Author: wyaadarsh
Email: aadarsh17i@gmail.com

Licensed under the MIT License, you may not use this file except in compliance with the License.
You may obtain a copy of the License at https://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under the License is
distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

See the License for the specific language governing permissions and limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wyaadarsh/fetch/core"
)

var downloadCmd = &cobra.Command{
	Use:   "download <URL> <filename>",
	Short: "Download a file from the internet",
	Long: `Download a file from the internet.
	Make Sure you have twice the space of the file to be downloaded in your disk.
	For example, if you are downloading a file of size 1GB, make sure you have 2GB of free space in your disk.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			os.Exit(0)
		}

		core.CreateSqliteDB()
		db := core.Get_DB()
		defer db.Close()
		var hist *core.History = new(core.History)

		location, err := cmd.Flags().GetString("path")
		if err != nil || location == "" {
			location = "$HOME/Downloads"
		}
		chunks, err := cmd.Flags().GetInt("threads")
		if err != nil || chunks == 0 {
			fmt.Println("Defaulting to 10 threads")
			chunks = 10
		}

	},
}
