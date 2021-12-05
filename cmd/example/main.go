package main

import (
	"flag"
	"io"
	"os"
	"strings"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	flagToComputeExpFromTerminal = flag.String("e", "", "Expression to compute")
	flagToReadExpFromFile = flag.String("f", "", "File with expression to compute")
	flagToWriteResultToFile = flag.String("o", "", "File with result of compute")
)

func main() {
	flag.Parse()
	var input io.Reader
	var output io.Writer

	if *flagToComputeExpFromTerminal != "" {
		input = strings.NewReader(*flagToComputeExpFromTerminal)
	}

	if *flagToReadExpFromFile != "" {
		dataFromFile, err := os.Open(*flagToReadExpFromFile)
		if err != nil {
			os.Stderr.WriteString("Error while reading expression from file!")
		}
		input = dataFromFile
	}

	if *flagToWriteResultToFile != "" {
		dataToFile, err := os.Create(*flagToWriteResultToFile)
		if err != nil {
			os.Stderr.WriteString("Error while writing result to file!")
		}
		output = dataToFile
	}

	if input == nil {
		os.Stderr.WriteString("No expression detected")
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}