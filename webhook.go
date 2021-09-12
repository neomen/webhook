package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var (
	workdir string
	secret  string
	port    string
)

func init() {
	//TODO check .ssh folder premitions
	workdir = os.Getenv("WEBHOOK_WORKDIR")
	secret = os.Getenv("WEBHOOK_SECRET")
	port = os.Getenv("WEBHOOK_PORT")
	if port == "" {
		port = "8000"
	}
	doPull()
}

// do git pull command
func doPull() (err error) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s; git pull;", workdir))
	log.Printf("run it: cd %s; git pull;", workdir)
	err = cmd.Run()
	return
}

// webhook heandler
func handle(w http.ResponseWriter, r *http.Request) {
	querySecret := r.URL.Query().Get("secret")
	if querySecret != "" && querySecret == secret {
		err := doPull()
		if err != nil {
			log.Printf("%s\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func main() {
	http.HandleFunc("/", handle)
	fmt.Printf("Webhook listener started at %s port\n", port)
	http.ListenAndServe(":"+port, nil)
}
