package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println("filename", filename)

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("error", err)
		return
	}

	str := string(content)

	length := len(str)

	if str[0] != '{' || str[length-1] != '}' {
		os.Exit(1)
	}

	fmt.Println(str)

}
