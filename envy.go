package envy

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// DefaultEnvFiles are the default env files used if none are provided to Load
// or MustLoad.
var DefaultEnvFiles = []string{
	".env",
	".env.local",
}

// Load loads the given env files. Later files overwrite earlier files.
func Load(envFiles ...string) error {
	if len(envFiles) == 0 {
		envFiles = DefaultEnvFiles
	}
	return loadFiles(envFiles)
}

// MustLoad loads the given env files. If a file is missing, it prints an error and exits.
func MustLoad(envFiles ...string) {
	err := Load(envFiles...)
	if err != nil {
		logFatalf("fatal error loading env files: %v", err)
	}
}

func loadFiles(envFiles []string) error {
	for _, f := range envFiles {
		bytes, err := os.ReadFile(f)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("error loading env file: %q: %w", f, err)
		}
		lines := strings.Split(string(bytes), "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				os.Setenv(parts[0], parts[1])
			}
		}
	}
	return nil
}

var logFatalf = func(format string, v ...any) {
	log.Fatalf(format, v...)
}
