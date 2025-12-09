package envy

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Load loads the given env files. Later files overwrite earlier files.
func Load(envFiles ...string) {
	loadFiles(envFiles, false)
}

// MustLoad loads the given env files. If a file is missing, it prints an error and exits.
func MustLoad(envFiles ...string) {
	loadFiles(envFiles, true)
}

func loadFiles(envFiles []string, exitOnMissing bool) {
	for _, f := range envFiles {
		bytes, err := os.ReadFile(f)
		if err != nil && (exitOnMissing || !errors.Is(err, os.ErrNotExist)) {
			fmt.Printf("fatal: unable to read %q: %v\n", f, err)
			exit(1)
		}
		lines := strings.Split(string(bytes), "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				os.Setenv(parts[0], parts[1])
			}
		}
	}
}

var exit = func(code int) {
	os.Exit(code)
}
