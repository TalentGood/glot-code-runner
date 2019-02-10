package perl6

import (
	"../../cmd"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	return cmd.RunStdin(workDir, stdin, maxTimeout, "perl6", files[0])
}
