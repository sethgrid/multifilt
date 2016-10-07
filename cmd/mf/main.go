package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/sethgrid/multifilt"
)

func init() {
	AppVersion := "1.0.0"
	flag.Usage = func() {
		fmt.Printf("Usage: %s (multifilter)\n\n", os.Args[0])
		fmt.Printf("Version %s, Compiled with %s\n\n", AppVersion, runtime.Version())
		fmt.Printf("Description\n")
		fmt.Printf("Filter out lines from an input source based on lines in a filter file.\n")
		fmt.Printf("A filter file of with two lines, 'ab' and 'bc', will filter out lines from the input source that contain either entry.\n\n")
		fmt.Printf("Examples\n")
		fmt.Printf("cat input | %s filter_file > output \n", os.Args[0])
		fmt.Printf("%s -in input -out ouput -f filter_file \n", os.Args[0])
		fmt.Printf("\n")
		flag.PrintDefaults()
	}
}

func main() {
	var fileIn, fileOut, fileFilter string
	var requireFullMatch bool
	flag.StringVar(&fileIn, "in", "", "file in, default stdin")
	flag.StringVar(&fileIn, "out", "", "file out, default stdout")
	flag.StringVar(&fileFilter, "f", "", "file filter, use -f or provide as single argument")
	flag.BoolVar(&requireFullMatch, "a", false, "filtered lines must match the whole line in the filter ('a' for match all)")
	flag.Parse()

	var in *os.File
	var out *os.File
	var filter *os.File
	var err error

	if fileIn == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(fileIn)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if fileOut == "" {
		out = os.Stdout
	} else {
		out, err = os.Open(fileOut)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if fileFilter == "" {
		if len(os.Args) >= 2 {
			filter, err = os.Open(os.Args[1])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		} else {
			fmt.Println("Provide a filter file. See -h")
			os.Exit(1)
		}
	} else {
		filter, err = os.Open(fileFilter)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	err = multifilt.Filter(in, filter, out, requireFullMatch)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
