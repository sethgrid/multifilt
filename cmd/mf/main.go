package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sethgrid/multifilt"
)

func init() {
	AppVersion := "1.2.3"
	flag.Usage = func() {
		fmt.Printf("Usage: %s (multifilter) [options] [argument]\n\n", os.Args[0])
		fmt.Printf("Version %s, Compiled with %s\n\n", AppVersion, runtime.Version())
		fmt.Printf("Description\n")
		fmt.Printf("Filter out lines from an input source based on lines in a filter file and/or -v flags.\n")
		fmt.Printf("A filter file with two lines, 'ab' and 'bc', will filter out lines from the input source that contain either entry.\n\n")
		fmt.Printf("Examples\n")
		fmt.Printf("cat input | %s -a filter_file > output \n", os.Args[0])
		fmt.Printf("cat input | %s -v foo -v bar -v raz filter_file > output \n", os.Args[0])
		fmt.Printf("%s -in input -out ouput -f filter_file \n", os.Args[0])
		fmt.Printf("\n")
		flag.PrintDefaults()
	}
}

// used for allowing multiple -v flags to stack into a slice of strings
type strSlice []string

func main() {
	var fileIn, fileOut, fileFilter string
	var filterList strSlice
	var requireFullMatch bool
	flag.StringVar(&fileIn, "in", "", "file in, default stdin")
	flag.StringVar(&fileOut, "out", "", "file out, default stdout")
	flag.StringVar(&fileFilter, "f", "", "file filter, use -f or provide as single argument")
	flag.Var(&filterList, "v", "specify multiple -v params to filter on each")
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
		out, err = os.OpenFile(fileOut, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil && err != os.ErrNotExist {
			fmt.Println(">", err.Error())
			os.Exit(1)
		}
	}

	var usingFlagFilters bool
	var usingFilterFile bool

	if fileFilter == "" {
		if len(os.Args) >= 2 {
			usingFlagFilters = os.Args[1] == "-v"

			// verify that the second to last element is not a flag key
			// verify that the last element is not a flag
			// if not a flag key, then last argument is filter file name
			if os.Args[len(os.Args)-2] != "-v" && !strings.HasPrefix(os.Args[len(os.Args)-1], "-") {
				usingFilterFile = true

				filter, err = os.Open(os.Args[len(os.Args)-1])
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			}
		} else if !usingFlagFilters {
			fmt.Println("Provide a filter file. See -h")
			os.Exit(1)
		}
	} else {
		usingFilterFile = true
		filter, err = os.Open(fileFilter)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if usingFlagFilters && usingFilterFile {
		flReader := strings.NewReader("\n" + strings.Join(filterList, "\n"))
		mr := io.MultiReader(filter, flReader)
		err = multifilt.Filter(in, mr, out, requireFullMatch)

	} else if usingFlagFilters {
		r := strings.NewReader(strings.Join(filterList, "\n"))
		err = multifilt.Filter(in, r, out, requireFullMatch)

	} else if usingFilterFile {
		err = multifilt.Filter(in, filter, out, requireFullMatch)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// String, the second method for flag.Value interface
func (s *strSlice) String() string {
	return strings.Join(*s, ",")
}

// Set, the second method for flag.Value interface
func (s *strSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}
