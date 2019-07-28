package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"github.com/greg-rychlewski/image-compare/csvutil"
	"github.com/greg-rychlewski/image-compare/fileutil"
)

// Build information from linker
var version, gitHash, buildTime, goBuildVersion string

// User-specified command-line flags
var isVersionFlagPresent bool
var inputPath, outputPath string

func init() {
	// Get default output path 
	defaultOutput := fileutil.GetDefaultOutputPath()

	// Initialize flag information
	flag.StringVar(&inputPath, "in", "", "Path to the input csv. (REQUIRED)")
	flag.StringVar(&outputPath,"out", defaultOutput, "Path to output csv. If not provided, a time-stamped file will be generated in the current directory.")
	flag.BoolVar(&isVersionFlagPresent, "version", false, "Print version information.")
}

func main() {
	// Process command-line flags
	flag.Parse()
	validateFlags()

	// Run main logic
	err := run()

	if err != nil {
		fmt.Fprintln(os.Stderr, "(ERROR)", err)
		os.Exit(1)
	}

}

func run() error {
	// Open input file
	inputFile, err := os.Open(inputPath)

	if err != nil {
		return err
	}

	defer inputFile.Close()

	// Create output file
	outputFile, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	// Process input file
	csvutil.Process(inputFile, outputFile)

	return nil
}

func validateFlags() {
	// If user specifies version flag, print version information and exit
	if isVersionFlagPresent {
		fmt.Printf("Version: %s (%s %s)\n\n", version, runtime.GOOS, runtime.GOARCH)
		fmt.Printf("For Developers Only:\n")
		fmt.Printf("  Built Under: %s\n", goBuildVersion)
		fmt.Printf("  UTC Build Time: %s\n", buildTime)
		fmt.Printf("  Git Hash: %s", gitHash)
		os.Exit(0)
	}

	// If user doesn't specify input path, exit
	if inputPath == "" {
		fmt.Printf("Usage of %s \n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

