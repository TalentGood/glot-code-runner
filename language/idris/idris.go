package idris

import (
	"../../cmd"
	"path/filepath"
)

func Run(files []string, maxTimeout int64,stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, maxTimeout,"idris", "-o", binName, files[0])
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, maxTimeout, binPath)
}
