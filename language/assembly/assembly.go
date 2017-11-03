package assembly

import (
	"github.com/prasmussen/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	objName := "a.o"
	binName := "a.out"

	stdout, stderr, err, elapsedTime, usedMemory := cmd.Run(workDir, "nasm", "-f", "elf64", "-o", objName, files[0])
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	stdout, stderr, err, elapsedTime, usedMemory = cmd.Run(workDir, "ld", "-o", binName, objName)
	if err != nil {
		return stdout, stderr, err, elapsedTime, usedMemory
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, binPath)
}
