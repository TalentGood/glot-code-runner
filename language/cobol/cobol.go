package cobol

import (
	"../../cmd"
	"../../util"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	sourceFiles := util.FilterByExtension(files, "cob")
	args := append([]string{"cobc", "-x", "-o", binName}, sourceFiles...)
	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, maxTimeout, args...)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, maxTimeout, binPath)
}
