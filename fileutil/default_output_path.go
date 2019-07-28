package fileutil

import (
	"os"
	"path"
	"github.com/greg-rychlewski/image-compare/timeutil"
)

// Create default output path if user doesn't specify one

func GenerateOutputPath()  (string, error) {
        currentDirectory, err := os.Getwd()

	if err != nil {
		return "", err
	}

	datetimeFormat := "20060102150405"

        return path.Join(currentDirectory,"out_" + timeutil.GetCurrentTime(datetimeFormat)), nil
}
