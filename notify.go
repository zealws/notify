package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const ADDR string = "127.0.0.1:42434"

type Notification struct {
	Title   string
	Content string
}

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

func sendNotification(prog string) {
	not := Notification{Title: prog, Content: "Command '" + strings.Join(os.Args[1:], " ") + "' finished."}
	buf := &bytes.Buffer{}

	err := json.NewEncoder(buf).Encode(not)
	if err != nil {
		fmt.Printf("Could not encode json: %v\n", err)
		os.Exit(1)
	}

	_, err = http.Post("http://"+ADDR, "application/json", buf)
	if err != nil {
		fmt.Printf("Could not send notification: %v\n", err)
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
	sendNotification(prog)
}
