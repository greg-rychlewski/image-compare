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
	// Run main logic
	err := run()

	if err != nil {
		fmt.Fprintln(os.Stderr, "(ERROR)", err)

		_, ok := err.(*flagutil.FlagError)

		if ok {
			fmt.Printf("Usage of %s \n", os.Args[0])
			flag.PrintDefaults()
		}

		os.Remove(outputPath)
		os.Exit(1)
	}

}

func run() error {
	// Parse and validate command-line flags
	flag.Parse()

	if isVersionFlagPresent {
		flagutil.PrintVersionInfo(version, goBuildVersion, buildTime, gitHash)

		return nil
	}

	if err := flagutil.ValidateInputPath(inputPath); err != nil {
		return err
	}

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
	numProcessedRows, err := csvutil.Process(inputFile, outputFile, !isNoHeaderFlagPresent)

	if err != nil {
		return err
	}

	fmt.Printf("%d image pairs successfully processed\n", numProcessedRows)
	fmt.Printf("Output saved to %s", outputPath)

	return nil
}
