package flagutil

import (
    "fmt"
    "runtime"
)

// Print app version + extra build information for developers
func PrintVersionInfo(version, goBuildVersion, buildTime, gitHash string) {
    fmt.Printf("\nImage Comparer Version %s (%s %s)\n\n", version, runtime.GOOS, runtime.GOARCH)
    fmt.Printf("For Developers Only:\n")
    fmt.Printf("  Built Under: %s\n", goBuildVersion)
    fmt.Printf("  UTC Build Time: %s\n", buildTime)
    fmt.Printf("  Git Hash: %s", gitHash)
}
