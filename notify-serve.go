package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

const ADDR string = "127.0.0.1:42434"

type Notification struct {
	Title   string
	Content string
}

func notify(not Notification) error {
	cmd := exec.Command("/usr/bin/notify-send", not.Title, not.Content)
	return cmd.Run()
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	var not Notification
	err := json.NewDecoder(r.Body).Decode(&not)
	if err != nil {
		fmt.Printf("Could not read JSON body: %v\n", err)
		return
	}
	err = notify(not)
	if err != nil {
		fmt.Printf("Could not send notification: %v\n", err)
		return
	}
}

func main() {
	s := &http.Server{
		Addr:         ADDR,
		Handler:      http.HandlerFunc(notifyHandler),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	err := s.ListenAndServe()
	panic(err)
}
