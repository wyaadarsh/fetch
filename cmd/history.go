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
	"time"

	"github.com/spf13/cobra"
	"github.com/wyaadarsh/fetch/core"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Fetch your download history",
	Long:  `Display the history of downloaded files using fetch`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetch Download History as of ", time.Now().Format("2006-01-02 15:04 Monday"))
		core.CreateSqliteDB()
		db := core.Get_DB()
		no_of_rows, err := cmd.Flags().GetInt("list")
		if err != nil {
			fmt.Printf("Defaulting to 10 rows\n")
			no_of_rows = 10
		}
		list := core.GetHistory(db, no_of_rows)
		for _, row := range list {
			fmt.Printf("%v\t%v\t%v\t%v\n", row.Success, row.FileName, row.FileSize, row.Date)
		}
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean your download history",
	Long:  `Delete all records of downloaded files using fetch`,
	Run: func(cmd *cobra.Command, args []string) {
		core.CreateSqliteDB()
		db := core.Get_DB()
		err := core.DeleteAllHistory(db)
		if err != nil {
			fmt.Println("Error while cleaning history", err)
		} else {
			fmt.Println("History cleaned successfully")
		}
	},
}

func init() {
	historyCmd.AddCommand(cleanCmd)
	historyCmd.PersistentFlags().Int("list", 10, "Number of rows to display")
	rootCmd.AddCommand(historyCmd)
}
