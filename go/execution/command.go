package execution

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func executeCommand(command string, args ...string) {
	// Command to execute
	cmd := exec.Command(command, args...)

	// Create pipes to capture stdout and stderr
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

	// Start the command
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read stdout
	stdout, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		fmt.Println("Error reading stdout:", err)
		return
	}

	// Read stderr
	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		fmt.Println("Error reading stderr:", err)
		return
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command finished with error:", err)
		return
	}

	// Print stdout and stderr
	fmt.Println("Stdout:", string(stdout))
	fmt.Println("Stderr:", string(stderr))
}