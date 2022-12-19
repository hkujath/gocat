package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Minimum of 1 argument is required!")
		os.Exit(1)
	}

	for _, fname := range os.Args[1:] {
		fmt.Printf("Content of [%s]\n", fname)

		// Open file
		f, err := os.Open(fname)
		if err != nil {
			log.Printf("Error while opening file: %s\n%s", fname, err)
			f.Close()
			continue
		}

		// Copy file to output
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			log.Printf("Error while printing output of %s\n%s", fname, err)
		}
		f.Close()

	}

}
