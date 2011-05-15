package main

import (
	"fmt"
	"flag"
	"os"
)

func cat(file *os.File) {
	nbyte := 512
	bytes := make([]byte, nbyte)
	for {
		if n, err := file.Read(bytes); err != nil {
			if err != os.EOF {
				fmt.Printf("Fail to read %s: %s", file.Name(), err)
			}
			return
		} else {
			os.Stdout.Write(bytes[:n])
		}
	}
}

func main() {
	flag.Parse()
	var file *os.File
	var err os.Error
	if flag.NArg() == 0 {
		file = os.Stdin
	} else {
		filename := flag.Arg(0)
		file, err = os.Open(filename)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}
	cat(file)
}
