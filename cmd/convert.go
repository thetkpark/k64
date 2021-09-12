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
	"github.com/thetkpark/k64/utils"
	"io/ioutil"
	"os"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

var filePath string
var outFilePath string
var isSave bool

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert string in the data fields to base64.",
	Long:  `Convert string in the data fields to base64. The output will print to the stdout by default.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Open file
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Unable to open file")
			os.Exit(1)
		}

		// Parse yaml
		kConfig, err := yaml.Parse(string(file))
		if err != nil {
			fmt.Println("Unable to parse yaml")
			os.Exit(1)
		}
		dataMap := kConfig.GetDataMap()

		// Check if the kind is secret
		if kind := kConfig.GetKind(); kind != `Secret` {
			var isContinue string
			fmt.Printf("%s is not a secret kind. Do you want to continue encoding? (Y)es, (N)o: ", filePath)
			_, err = fmt.Scan(&isContinue)
			if err != nil {
				fmt.Println("Unable to get answer")
				os.Exit(1)
			}
			if isContinue == "N" || isContinue == "n" {
				os.Exit(0)
			}
		}

		// Load data fields and convert to base64
		for key, value := range dataMap {
			dataMap[key] = utils.ToBase64(value)
		}
		kConfig.SetDataMap(dataMap)
		strConfig, err := kConfig.String()
		if err != nil {
			fmt.Println("Unable to get string from config")
			os.Exit(1)
		}

		// Output or save
		if !isSave && len(outFilePath) == 0 {
			fmt.Println(strConfig)
		} else {
			if len(outFilePath) == 0 {
				outFilePath = filePath
			}
			if err := ioutil.WriteFile(outFilePath, []byte(strConfig), 0644); err != nil {
				fmt.Println("Unable write back to", filePath)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&filePath, "file", "f", "", "File that you want to convert secret string to base64 string")
	convertCmd.Flags().StringVarP(&outFilePath, "out", "o", "", "Write the output to this file path")
	convertCmd.Flags().BoolVarP(&isSave, "save", "s", false, "Save the output to the same file")
}
