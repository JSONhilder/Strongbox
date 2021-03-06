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
	"github.com/JSONhilder/strongbox/internal/database"
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save new entry to strongbox accounts",
	Long: `Save a new entry with given details
	Can also generate passwords by using gen=passwordLength (gen=25)
	`,
	Example: `
	strongbox save <Account Name> <username> <password> <optional url>

	Strongbox can generate a password for you:
	strongbox save <Account Name> <username> gen=30 <optional url>
	
	`,
	Args: cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		newAccount := database.Account{
			Name:     args[0],
			Username: args[1],
			Password: args[2],
			Url:      "null",
		}

		if len(args) == 4 {
			newAccount.Url = args[3]
		}

		database.CreateAccount(newAccount)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
