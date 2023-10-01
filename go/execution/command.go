package execution

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"unicode/utf8"
)

// func test(cmd *exec.Cmd) {
// 	time.Sleep(4 * time.Second)
// 	cmd.Process.Signal(syscall.SIGTERM)
// }

func captureStandardOutOrErrAndShow(stdIOutOrErrPipe io.ReadCloser, stdChan chan bool) {
	oneRuneStdOutOrErr := make([]byte, utf8.UTFMax)
	for {
		count, err := stdIOutOrErrPipe.Read(oneRuneStdOutOrErr)
		if err != nil {
			break
		}
		fmt.Printf("%s", oneRuneStdOutOrErr[:count])
	}
	stdChan <- true
}

func executeCommand(command string, envVars map[string]string, args ...string) {
	// Command to execute
	var err error
	stdOutChan := make(chan bool, 1)
	stdErrChan := make(chan bool, 1)
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
	for key, value := range envVars {
		cmd.Env = append(cmd.Env, key+"="+value)
	}
	// Used if not record to be stored, directly piping to stream
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// can be Used to store data before piping to prompt
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	go captureStandardOutOrErrAndShow(stdoutPipe, stdOutChan)
	go captureStandardOutOrErrAndShow(stderrPipe, stdErrChan)
	// oneRuneStdout := make([]byte, utf8.UTFMax)
	// for {
	// 	count, err := stdoutPipe.Read(oneRuneStdout)
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Printf("%s", oneRuneStdout[:count])
	// }

	// oneRuneStderr := make([]byte, utf8.UTFMax)
	// for {
	// 	count, err := stderrPipe.Read(oneRuneStderr)
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Printf("%s", oneRuneStderr[:count])
	// }

	// // Create pipes to capture stdout and stderr
	// stdoutPipe, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// stderrPipe, err := cmd.StderrPipe()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// // Start the command
	// err = cmd.Start()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// // Read stdout
	// stdout, err := ioutil.ReadAll(stdoutPipe)
	// if err != nil {
	// 	fmt.Println("Error reading stdout:", err)
	// 	return
	// }

	// // Read stderr
	// stderr, err := ioutil.ReadAll(stderrPipe)
	// if err != nil {
	// 	fmt.Println("Error reading stderr:", err)
	// 	return
	// }
	// go test(cmd)
	// combinedOutput, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// Wait for the command to finish
	go func() {
		<-osSignal
		cmd.Process.Kill()
	}()
	// test(cmd)
	<-stdOutChan
	<-stdErrChan

	err = cmd.Wait()
	if err != nil && err.Error() != "signal: killed" {
		fmt.Println("Command finished with error:", err)
		return
	}

	// fmt.Println("Command finished")
	// Print stdout and stderr
	// fmt.Println(string(combinedOutput))
	// fmt.Println("Stdout:", string(stdout))
	// fmt.Println("Stderr:", string(stderr))
}
