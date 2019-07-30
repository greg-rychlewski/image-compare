package main

import (
	"github.com/greg-rychlewski/image-compare/csvutil"
	"github.com/greg-rychlewski/image-compare/flagutil"
	"flag"
	"fmt"
	"os"
)

// Build information from build script
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

		os.Exit(1)
        }

        // Process input file
	fmt.Println("processing csv file...")

        numProcessedPairs, err := csvutil.Process(inputPath, outputPath, !isNoHeaderFlagPresent)

        if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Remove(outputPath)

		os.Exit(1)
        }

        fmt.Printf("%d image pairs successfully processed\n", numProcessedPairs)
        fmt.Printf("output saved to %s", outputPath)
}
