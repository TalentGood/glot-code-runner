package fsharp

import (
	"../../cmd"
	"../../util"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	binName := "a.exe"

	sourceFiles := reverse(util.FilterByExtension(files, "fs"))
	args := append([]string{"fsharpc", "--out:" + binName}, sourceFiles...)
	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, maxTimeout, args...)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, maxTimeout, "mono", binPath)
}

func reverse(files []string) []string {
	reversed := make([]string, 0, len(files))

	for i := len(files) - 1; i >= 0; i-- {
		reversed = append(reversed, files[i])
	}

	return reversed
}
