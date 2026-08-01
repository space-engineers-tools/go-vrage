package vrage

import (
	"runtime/debug"
	"strings"
)

// Repository is the module path for the go-vrage package
const Repository string = "github.com/space-engineers-tools/go-vrage"

// Version is the current version of the go-vrage package
var Version = func() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	for _, dep := range info.Deps {
		if dep.Path == Repository {
			// If it's a local development build without a tag,
			// Go might show a pseudo-version. You can clean it up or keep it as-is.
			return strings.TrimPrefix(dep.Version, "v")
		}
	}

	// Fallback if the module path isn't found in dependencies
	// (e.g., during local package testing within the same module)
	return "dev"
}()
