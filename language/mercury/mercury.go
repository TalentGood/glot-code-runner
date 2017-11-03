package mercury

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"github.com/prasmussen/glot-code-runner/util"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	sourceFiles := util.FilterByExtension(files, "m")
	args := append([]string{"mmc", "-o", binName}, sourceFiles...)
	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, binPath)
}
