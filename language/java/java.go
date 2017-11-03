package java

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"path/filepath"
	"fmt"
)

func Run(files []string, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	fname := filepath.Base(files[0])

	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, "javac", fname)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}
	fmt.Println("PASSSSSSSSSED")

	return cmd.RunStdin(workDir, stdin, "java", className(fname))
}

func className(fname string) string {
	ext := filepath.Ext(fname)
	return fname[0 : len(fname)-len(ext)]
}
