package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
	"encoding/csv"
)

// Build information from linker
var version, gitHash, buildTime, goBuildVersion string

// User-specified command-line flags
var isVersionFlagPresent bool
var inputPath, outputPath string

func init() {
	// Initialize flag information
	flag.StringVar(&inputPath, "in", "", "Path to the input csv. (REQUIRED)")
	flag.StringVar(&outputPath,"out", "out_" + getCurrentTime(), "Path to output csv. If this is not specified, a time-stamped csv will be saved to the current directory.")
	flag.BoolVar(&isVersionFlagPresent, "version", false, "Print version information.")
}

func main() {
	// Process command-line flags
	flag.Parse()
	validateFlags()

	// Open input file
	inputFile, err := os.Open(inputPath)
	checkBreakingError(err)
	defer inputFile.Close()

	// Create output file
	outputFile, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
	checkBreakingError(err)
	defer outputFile.Close()

	// Process input file
	parseCSV(inputFile, outputFile)
}

func parseCSV(inputFile *os.File, outputFile *os.File) {
	csvReader := csv.NewReader(inputFile)
	csvWriter := csv.NewWriter(outputFile)


	// Parse csv one row at a time, treating the header separately
	isHeader := true

	for {
		row, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		checkBreakingError(err)

		if !isHeader {
			row = append(row, "0.1", "0.001")
		} else {
			row = append(row, "similar", "elapsed")
		}

		csvWriter.Write(row)
		csvWriter.Flush()

		if isHeader {
			isHeader = false
		}
	}
}

func getCurrentTime() string{
	// Returns time in YYYYMMDDhhmmss format
	return time.Now().Format("20060102150405")
}

func checkBreakingError(err error) {
	if err != nil {
		fmt.Println("(ERROR)", err)
		os.Exit(1)
	}
}

func validateFlags() {
	// If user specifies version flag, print version information and exit
	if isVersionFlagPresent {
		fmt.Printf("Version: %s (%s %s)\n\n", version, runtime.GOOS, runtime.GOARCH)
		fmt.Printf("For Developers Only:\n\n")
		fmt.Printf("Built Under: %s\n", goBuildVersion)
		fmt.Printf("UTC Build Time: %s\n", buildTime)
		fmt.Printf("Git Hash: %s", gitHash)
		os.Exit(0)
	}

	// If user doesn't specify input path, exit
	if inputPath == "" {
		fmt.Printf("Usage of %s \n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

