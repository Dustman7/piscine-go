package main

import (
	"fmt"
	"os"
)

func main() {
	as := os.Args[1:]
	if len(as) >= 1 {
	} else {
		fmt.Println("File name missing")
		return
	}

	if len(as) == 1 {
		content, err := os.Open(as[0])
		if err != nil {
			fmt.Println(err)
		} else {
			data := make([]byte, 14)
			content.Read(data)
			fmt.Println(string(data))
			content.Close()
		}
	} else {
		fmt.Println("Too many arguments")
	}
}
