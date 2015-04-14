package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const DEFAULT_ADDR string = "0.0.0.0:42434"

func notify(not string) error {
	cmd := exec.Command("/usr/bin/notify-send", "Notify", not)
	return cmd.Run()
}

func notifyHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		fmt.Printf("Could not read notification: %v\n", err)
		return
	}
	err = notify(buf.String())
	if err != nil {
		fmt.Printf("Could not send notification: %v\n", err)
		return
	}
}

func getAddr() string {
	addr := os.Getenv("NOTIFY_ADDR")
	if addr != "" {
		return addr
	}
	return DEFAULT_ADDR
}

func main() {
	addr := getAddr()
	s := &http.Server{
		Addr:         addr,
		Handler:      http.HandlerFunc(notifyHandler),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	fmt.Println("Listening on", addr)
	err := s.ListenAndServe()
	panic(err)
}
