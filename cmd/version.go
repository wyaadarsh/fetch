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

	"github.com/spf13/cobra"
)

var version string = "1.0.0"

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Display the version of Fetch",
	Long:  `Display the version of Fetch`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(" Fetch version " + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
