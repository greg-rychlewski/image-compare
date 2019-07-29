package flagutil

import (
	"flag"
	"github.com/greg-rychlewski/image-compare/fileutil"
)

// Initialize command-line flag information

func InitFlags(inputPathPtr *string, outputPathPtr *string, isVersionFlagPresentPtr *bool) {
	// Get default output path 
	defaultOutput := fileutil.GetDefaultOutputPath()

	// Initialize flag information
	flag.StringVar(inputPathPtr, "in", "", "Path to the input csv. (REQUIRED)")
	flag.StringVar(outputPathPtr,"out", defaultOutput, "Path to output csv. If not provided, a time-stamped file will be generated in the current directory.")
	flag.BoolVar(isVersionFlagPresentPtr, "version", false, "Print version information.")
}
