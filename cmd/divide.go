/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// divideCmd represents the divide command
var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divide given IP into given divisor",
	Long: `Takes a base IP and a divisor and returns subnets in CIDR notation.
	For example: 

		APPNAME COMMAND ARG ARG --FLAG`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("divide called")
		if viper.IsSet("address") {
			fmt.Println("address set: ", viper.GetString("address"))
		}
		if viper.IsSet("divisor") {
			fmt.Println("divisor set: ", viper.GetString("divisor"))
		}
	},
}

var address string
var divisor int

func init() {
	rootCmd.AddCommand(divideCmd)

	// Define command flags and config settings

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Local flag for command:
	// add StringVarP flag for command
	divideCmd.Flags().StringVarP(&address, "address", "a", "", "The base IP address to divide.")
	divideCmd.Flags().IntVarP(&divisor, "divisor", "d", 1, "The amount of subnets to split network into.")
	// bind Cobra pflag defined above with Viper
	viper.BindPFlag("address", divideCmd.Flags().Lookup("address"))
	viper.BindPFlag("divisor", divideCmd.Flags().Lookup("divisor"))
}
