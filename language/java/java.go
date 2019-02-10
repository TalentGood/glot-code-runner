package java

import (
	"../../cmd"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	fname := filepath.Base(files[0])

	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, maxTimeout, "javac", fname)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}
	return cmd.RunStdin(workDir, stdin, maxTimeout, "java", className(fname))
}

func className(fname string) string {
	ext := filepath.Ext(fname)
	return fname[0 : len(fname)-len(ext)]
}
