package main

import (
	"github.com/greg-rychlewski/image-compare/csvutil"
	"github.com/greg-rychlewski/image-compare/flagutil"
	"flag"
	"fmt"
	"os"
)

// Build variables imported from ldflags
var version, gitHash, buildTime, goBuildVersion string

// Command-line flags
var isVersionFlagPresent, isNoHeaderFlagPresent bool
var inputPath, outputPath string

func init() {
	// Initialize command-line flag information
	flagutil.InitFlags(&inputPath, &outputPath, &isVersionFlagPresent, &isNoHeaderFlagPresent)
}

func main() {
	// Parse and validate command-line flags
        flag.Parse()

        if isVersionFlagPresent {
                flagutil.PrintVersionInfo(version, goBuildVersion, buildTime, gitHash)

                return
        }

        if err := flagutil.ValidateInputPath(inputPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()

		return
        }

        // Process input file
	fmt.Printf("Processing %s...\n", inputPath)

        numProcessedPairs, err := csvutil.Process(inputPath, outputPath, !isNoHeaderFlagPresent)

        if err != nil {
		if csvErr, ok := err.(*csvutil.CsvError); ok {
			fmt.Fprintf(os.Stderr, "%s line %d). %s\n", inputPath, csvErr.Row, csvErr.Message)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Fprintln(os.Stderr, "Fatal error. Program exiting unsuccessfully.")
		os.Remove(outputPath)
		os.Exit(1)
        }

        fmt.Printf("%d image pairs successfully processed\n", numProcessedPairs)
        fmt.Printf("Output saved to %s", outputPath)
}
