package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	parseArgsPassed()
}

func parseArgsPassed() {
	argsPassed := os.Args
	// First argument will be the program name .Second would be filename. Eg: dummyFile is the file name in the below command
	//  go main.go dummyFile
	if len(argsPassed) > 1 {
		openFileByName(argsPassed[1])
	} else {
		fmt.Println("Error: Filename not provided")
		os.Exit(1)
	}
}

func openFileByName(fn string) {
	// This could also be done using ioutil.ReadFile but the assignment expects us to use os.Open
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, f)
	}
}
