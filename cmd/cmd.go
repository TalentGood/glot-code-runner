package cmd

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
	"syscall"
)

//type TimeLimitError string
//
//func (e TimeLimitError) Error() string {
//	return string(e)
//}

func Run(workDir string, args ...string) (string, string, error, int64, int64) {
	return RunStdin(workDir, "", args...)
}

func RunStdin(workDir, stdin string, args ...string) (string, string, error, int64, int64) {

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = workDir
	cmd.Stdin = strings.NewReader(stdin)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	start := time.Now()
	err := cmd.Start()

	var timer *time.Timer

	timer = time.AfterFunc(1 * time.Second, func() {
		timer.Stop()
		cmd.Process.Signal(syscall.SIGKILL)
		var message = "Time limit exceeded"
		stderr.Grow(len(message))
		stderr.Write([]byte(message))
		stdout.Reset();
	})

	cmd.Wait()

	elapsedTime := int64(time.Since(start))

	var usedMemory int64;

	if cmd.ProcessState != nil {
		usedMemory = cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss
	}

	if len(stderr.String()) != 0 {
		stdout.Reset();
	}

	if len(stdout.String()) > 9999999 {
		var message = "Output limit exceeded"
		stderr.Grow(len(message))
		stderr.Write([]byte(message))
		stdout.Reset();
	}

	return stdout.String(), stderr.String(), err, elapsedTime, usedMemory
}

func RunBash(workDir, command string) (string, string, error, int64, int64) {
	return Run(workDir, "bash", "-c", command)
}

func RunBashStdin(workDir, command, stdin string) (string, string, error, int64, int64) {
	return RunStdin(workDir, stdin, "bash", "-c", command)
}
