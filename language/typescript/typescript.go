package typescript

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"github.com/prasmussen/glot-code-runner/util"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	jsName := "a.js"

	// Find all typescript files and build compile command
	sourceFiles := util.FilterByExtension(files, "ts")
	args := append([]string{"tsc", "-out", jsName}, sourceFiles...)

	// Compile to javascript
	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	return cmd.RunStdin(workDir, stdin, "node", jsName)
}
