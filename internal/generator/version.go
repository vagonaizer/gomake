package generator

import (
	"fmt"
	"runtime"
)

var (
	// Version is set during build time
	Version = "dev"
	// GitCommit is set during build time
	GitCommit = "unknown"
	// BuildDate is set during build time
	BuildDate = "unknown"
)

// GetVersionInfo returns formatted version information
func GetVersionInfo() string {
	return fmt.Sprintf(`gomake version %s
Git commit: %s
Build date: %s
Go version: %s
OS/Arch: %s/%s`,
		Version,
		GitCommit,
		BuildDate,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH)
}
