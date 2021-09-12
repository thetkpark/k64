package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func OpenFile(filePath string) string {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Unable to open file", filePath)
		os.Exit(1)
	}
	return string(file)
}

func WriteToFile(filePath string, content string) {
	if err := ioutil.WriteFile(filePath, []byte(content), 0644); err != nil {
		fmt.Println("Unable write back to", filePath)
		os.Exit(1)
	}
}
