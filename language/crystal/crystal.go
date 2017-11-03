package crystal

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	return cmd.RunStdin(workDir, stdin, "crystal", "run", files[0])
}