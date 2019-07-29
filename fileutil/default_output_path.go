package fileutil

import (
	"github.com/greg-rychlewski/image-compare/timeutil"
	"path"
)

// Create default output path if user doesn't specify one

func GetDefaultOutputPath() string {
	datetimeFormat := "20060102150405"

        return path.Join("out_" + timeutil.GetCurrentTime(datetimeFormat))
}
