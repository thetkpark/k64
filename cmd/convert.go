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
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)


var filePath string
var isSave bool

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert string in the data fields to base64.",
	Long: `Convert string in the data fields to base64. The output will print to the stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Unable to open file")
			os.Exit(1)
		}

		kConfig, err := yaml.Parse(string(file))
		if err != nil {
			fmt.Println("Unable to parse yaml")
			os.Exit(1)
		}
		dataMap := kConfig.GetDataMap()

		for key, value := range dataMap {
			dataMap[key] = toBase64([]byte(value))
		}
		kConfig.SetDataMap(dataMap)
		strConfig, err := kConfig.String()
		if err != nil {
			fmt.Println("Unable to get string from config")
			os.Exit(1)
		}

		if !isSave {
			fmt.Println(strConfig)
			return
		}
		if err := ioutil.WriteFile(filePath, []byte(strConfig), 0644); err != nil {
			fmt.Println("Unable write back to", filePath)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&filePath, "file", "f", "", "File that you want to convert secret string to base64 string")
	convertCmd.Flags().BoolVarP(&isSave, "save", "s", false, "Save the output to the same file")
}

func toBase64(text []byte) string {
	return base64.StdEncoding.EncodeToString(text)
}

