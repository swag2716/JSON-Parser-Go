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

	if length < 2 {
		fmt.Println("Length = 0")
		os.Exit(1)
	}

	if str[0] != '{' || str[length-1] != '}' {
		fmt.Println("curly braces error")
		os.Exit(1)
	}
	flag := false
	for i := 1; i < length-1; i++ {
		k := i
		if str[k] != '"' {
			fmt.Println("syntax error")
			os.Exit(1)
		}
		for str[k] != '"' {
			k++
		}
		if flag && str[k] != ',' {
			fmt.Println("syntax error")
			os.Exit(1)
		}
		if !flag && str[k] != ':' {
			fmt.Println("syntax error")
			os.Exit(1)
		}
		flag = !flag
		i = k
	}

	fmt.Println(str)

}
