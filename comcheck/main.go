package main

import (
	"fmt"
	"os"
)

func main() {
	os := os.Args[1:]
	for t := range os {
		if os[t] == "01" || os[t] == "galaxy" || os[t] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
