package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// flag has the Usage() function, but it does
// not allow you to define positional arguments, so
// so this has to replace it.
func usage() {
	fmt.Fprintf(os.Stderr, "filesize returns human readable filesizes.\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\t filesize [-h] [-u <MB|KB|GB>] <file-names>\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(os.Stderr, "\t-%s %v\n", f.Name, f.Usage)
	})
}

func display_file_sizes(byte_divisor int, unit_name string) {
	for _, file := range flag.Args() {
		info, err := os.Stat(file)
		if err == nil {
			fmt.Fprintf(os.Stdout, "%s %.2f%s\n", file, float64(info.Size())/float64(byte_divisor), unit_name)

		}
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintf(os.Stderr, "WARNING: %s does not exist\n", file)
		}
	}

}

func main() {
	unit := flag.String("u", "KB", "The unit you want the file size displayed in. Supports KB, MB, GB. Defaults to KB")
	help := flag.Bool("h", false, "Display this help message")
	flag.Parse()
	flag.Usage = usage

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "ERROR: You must specify a file!\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// There are such things as kilo-bits, but
	// No one actually looks at files that way.
	// We assume people are always looking at bytes
	unit_lower := strings.ToLower(*unit)

	switch unit_lower {
	case "kb":
		display_file_sizes(1024, "KB")
	case "mb":
		display_file_sizes(1048576, "MB")
	case "gb":
		display_file_sizes(1073741824, "GB")
	default:
		fmt.Fprintf(os.Stderr, "ERROR: Unknown unit: %s! \n", *unit)
		flag.Usage()
		os.Exit(1)
	}
	// Not necessarily needed, but for consistency.
	os.Exit(0)
}
