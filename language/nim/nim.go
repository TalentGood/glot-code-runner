package nim

import (
	"../../cmd"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	return cmd.RunStdin(workDir, stdin, maxTimeout, "nim", "--hints:off", "--verbosity:0", "compile", "--run", files[0])
}
