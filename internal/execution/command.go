/*
Copyright 2023 Jasti Sri Radhe Shyam

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

	// Prepare standard output and error streams
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

	// Start the sub process
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Handle output in async fashion
	go captureStandardOutOrErrAndShow(stdoutPipe, stdOutChan)
	go captureStandardOutOrErrAndShow(stderrPipe, stdErrChan)

	// Wait for the command to finish or terminate the sub process if main process got terminate signal
	go func() {
		<-osSignal
		cmd.Process.Kill()
	}()
	<-stdOutChan
	<-stdErrChan

	err = cmd.Wait()
	if err != nil && err.Error() != "signal: killed" {
		fmt.Println("Command finished with error:", err)
		return
	}
}
