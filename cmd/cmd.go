package cmd

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
	"syscall"
	"fmt"
)

func Run(workDir string, args ...string) (string, string, error, int64, int64) {
	return RunStdin(workDir, "", args...)
}

func RunStdin(workDir, stdin string, args ...string) (string, string, error, int64, int64) {

	var stdout bytes.Buffer
	var stderr bytes.Buffer


	// Limit execution time in seconds
	var rTimeLimit syscall.Rlimit

	rTimeLimit.Max = 3
	rTimeLimit.Cur = 3

	err := syscall.Setrlimit(syscall.RLIMIT_CPU, &rTimeLimit)

	if err != nil {
		fmt.Println("Error Setting Rlimit ", err)
	}

	// Limit allocated memory
	//var rMemoryLimit syscall.Rlimit
	//
	//rMemoryLimit.Max = 80000
	//rMemoryLimit.Cur = 80000
	//
	//err = syscall.Setrlimit(0x9, &rMemoryLimit)
	//
	//if err != nil {
	//	fmt.Println("Error Setting Rlimit ", err)
	//}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = workDir
	cmd.Stdin = strings.NewReader(stdin)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	start := time.Now()
	cmd.Wait()
	elapsedTime := int64(time.Since(start))

	var usedMemory int64;

	if cmd.ProcessState != nil {
		usedMemory = cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss
	}

	return stdout.String(), stderr.String(), err, elapsedTime, usedMemory
}

func RunBash(workDir, command string) (string, string, error, int64, int64) {
	return Run(workDir, "bash", "-c", command)
}

func RunBashStdin(workDir, command, stdin string) (string, string, error, int64, int64) {
	return RunStdin(workDir, stdin, "bash", "-c", command)
}
