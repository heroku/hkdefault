package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	path, err := exec.LookPath("heroku")
	if err != nil {
		log.Fatal(err)
	}
	syscall.Exec(path, append([]string{"heroku"}, os.Args...), os.Environ())
}
