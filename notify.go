package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const DEFAULT_ADDR string = "127.0.0.1:42434"

func splitArgs() (prog string, args []string) {
	prog, args = os.Args[1], []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	return
}

func runCommand(prog string, args []string) {
	cmd := exec.Command(prog, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Could not run command: %v\n", err)
		os.Exit(1)
	}
}

func getAddr() string {
	addr := os.Getenv("NOTIFY_ADDR")
	if addr != "" {
		return addr
	}
	return DEFAULT_ADDR
}

func sendNotification() {
	not := "Command '" + strings.Join(os.Args[1:], " ") + "' finished."
	buf := &bytes.Buffer{}
	fmt.Fprint(buf, not)

	// Try sending through CLI first
	cmd := exec.Command("/usr/bin/notify-send", "Notify", not)
	err := cmd.Run()
	if err == nil {
		return
	}

	// Try sending through API only if CLI fails
	_, err = http.Post("http://"+getAddr(), "application/json", buf)
	if err != nil {
		fmt.Printf("Could not send notification: %v", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: notify CMD")
		os.Exit(1)
	}
	prog, args := splitArgs()
	runCommand(prog, args)
	sendNotification()
}
