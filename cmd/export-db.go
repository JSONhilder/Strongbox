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

// exportDbCmd represents the exportDb command
var exportDbCmd = &cobra.Command{
	Use:     "export-db",
	Short:   "Exported encrypted db file.",
	Long:    `Exported encrypted db file to provided absolute path.`,
	Args:    cobra.MinimumNArgs(1),
	Example: "strongbox export-db C:\\Path\\To\\Directory\\",
	Run: func(cmd *cobra.Command, args []string) {
		database.ExportDb(args[0])
	},
}

func init() {
	rootCmd.AddCommand(exportDbCmd)
}
