package kotlin

import (
	"../../cmd"
	"path/filepath"
	"strings"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	fname := filepath.Base(files[0])

	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, maxTimeout, "kotlinc", fname)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	return cmd.RunStdin(workDir, stdin, maxTimeout, "kotlin", className(fname))
}

func className(fname string) string {
	if len(fname) < 5 {
		return fname
	}

	ext := filepath.Ext(fname)
	name := fname[0 : len(fname)-len(ext)]
	return strings.ToUpper(string(name[0])) + name[1:] + "Kt"
}
