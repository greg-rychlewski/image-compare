package flagutil

import (
	"github.com/greg-rychlewski/image-compare/fileutil"
	"flag"
)

// Initialize command-line flag information
func InitFlags(inputPathPtr *string, outputPathPtr *string, isVersionFlagPresentPtr *bool, isNoHeaderFlagPresentPtr *bool) {
	defaultOutputPath := fileutil.GetDefaultOutputPath()

	flag.StringVar(inputPathPtr, "in", "", "Path to the input csv. (REQUIRED)")
	flag.StringVar(outputPathPtr,"out", defaultOutputPath, "Path to output csv. If not provided, a time-stamped file will be generated in the current directory.")
	flag.BoolVar(isVersionFlagPresentPtr, "version", false, "Print version information.")
	flag.BoolVar(isNoHeaderFlagPresentPtr, "no-header", false, "Specify this flag to indicate that the input csv file does not have a header.")
}
