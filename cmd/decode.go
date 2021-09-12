/*
Copyright © 2021 Sethanant Pipatpakorn <sethanant.p@icloud.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/thetkpark/k64/utils"
	"os"
	"sigs.k8s.io/kustomize/kyaml/yaml"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode data fields of secret from base64 to string",
	Long:  `Decode data fields of secret from base64 to string. The output will be printed to stdout by default`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires input file argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Open file
		filePath := args[0]
		file := utils.OpenFile(filePath)

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
			dataMap[key], err = utils.FromBase64(value)
			if err != nil {
				fmt.Println("Unable to decode base64 of key", key)
				os.Exit(1)
			}
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
			utils.WriteToFile(outFilePath, strConfig)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().StringVarP(&outFilePath, "out", "o", "", "Write the output to this file path")
	decodeCmd.Flags().BoolVarP(&isSave, "save", "s", false, "Save the output to the same file")
}
