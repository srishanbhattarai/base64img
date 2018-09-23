package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please provide the path to the image")
		os.Exit(1)
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read or open file: %v\n", err)
		os.Exit(1)

	}

	enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	enc.Write(b)
	enc.Close()
	fmt.Println()
}
