package elm

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])

	// Move bootstrap files into work dir
	stdout, stderr, err, elapsedTime, usedMemory := cmd.RunBash(workDir, "cp -rf /bootstrap/* .")
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	// Compile elm to javascript
	stdout, stderr, err, elapsedTime, usedMemory = cmd.Run(workDir, "elm-make", files[0], "--output", "elm.js")
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	// Run javascript with node via app.js from bootstrap
	return cmd.RunStdin(workDir, stdin, "node", "app.js")
}
