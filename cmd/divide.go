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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// divideCmd represents the divide command
var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divide given IP into given divisor",
	Long: `Takes a network IP address and a divisor and returns subnets in CIDR notation.
	For example: 

		APPNAME COMMAND -a ADDRESS -d DIVISOR`,
	Run: func(cmd *cobra.Command, args []string) {

		if !viper.IsSet("address") {
			panic("Something went horribly wrong...")
		}

		// max subnets: 2^(32-routing_prefix)
		// max addresses: 2^(32-routing_prefix) - 2
		fmt.Println("address set: ", viper.GetString("address"))
		fmt.Println("divisor set: ", viper.GetString("divisor"))

		address := viper.GetString("address")

		// check if standard 3 dots are in string
		if strings.Count(address, ".") != 3 {
			panic("That doesn't look like any IP address I've seen...")
		}

		// split IPv4 string into its 8 bit sections
		address_parts := strings.Split(address, ".")
		address_binary := []string{}

		// convert each section to binary
		for _, part := range address_parts {
			val, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				panic("That's no int...")
			}
			base2 := strconv.FormatInt(int64(val), 2)
			// need to append 0s because base 2 int not always octet
			for i := len(base2); i < 8; i++ {
				base2 = "0" + base2
			}
			address_binary = append(address_binary, base2)
		}
		fmt.Println(address_binary)
		// output slash cidr notation for each possible subnet mask?
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
