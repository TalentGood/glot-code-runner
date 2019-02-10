package erlang

import (
	"../../cmd"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])

	// Compile all files except the first
	for _, file := range files[1:] {
		stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, maxTimeout, "erlc", file)
		if err != nil {
			return stdout, stderr, err, elapsedTime, usedMemory
		}
	}

	// Run first file with escript
	return cmd.RunStdin(workDir, stdin, maxTimeout, "escript", files[0])
}
