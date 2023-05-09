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

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch is a CLI tool to download files from the internet",
	Long: `fetch is a CLI tool to download files from the internet.
It is a concurrent downloader that downloads files in parallel.
It is written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		start_time := time.Now()
		info := core.Make_Info(args[0], "$HOME/Downloads"+"/"+args[1], 10)
		err := core.Download(info)
		if err != nil {
			fmt.Printf("Encountered error: %v\n", err)
		} else {
			fmt.Printf("Download completed successfully in %f seconds!", time.Since(start_time).Seconds())
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
