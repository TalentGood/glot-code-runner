package scala

import (
	"../../cmd"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])

	args := append([]string{"scalac"}, files...)
	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, maxTimeout, args...)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	return cmd.RunStdin(workDir, stdin, maxTimeout, "scala", "Main")
}
