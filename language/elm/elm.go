package elm

import (
	"../../cmd"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])

	// Move bootstrap files into work dir
	stdout, stderr, err, elapsedTime, usedMemory := cmd.RunBash(workDir, maxTimeout, "cp -rf /bootstrap/* .")
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	// Compile elm to javascript
	stdout, stderr, err, elapsedTime, usedMemory = cmd.Run(workDir, maxTimeout, "elm-make", files[0], "--output", "elm.js")
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	// Run javascript with node via app.js from bootstrap
	return cmd.RunStdin(workDir, stdin, maxTimeout, "node", "app.js")
}
