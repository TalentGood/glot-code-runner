package clojure

import (
	"../../cmd"
	"path/filepath"
)

func Run(files []string, maxTimeout int64, stdin string) (string, string, error, int64, int64) {
	workDir := filepath.Dir(files[0])
	return cmd.RunStdin(workDir, stdin, maxTimeout, "java", "-cp", "/usr/share/java/clojure.jar", "clojure.main", files[0])
}
