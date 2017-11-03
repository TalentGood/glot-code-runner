package scala

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])

	args := append([]string{"scalac"}, files...)
	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	return cmd.RunStdin(workDir, stdin, "scala", "Main")
}
