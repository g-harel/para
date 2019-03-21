package main

import (
	"os/exec"
)

// TODO configurable
const name = "sh"

var printer = &Printer{}

func main() {
	_, err := exec.LookPath(name)
	if err != nil {
		printer.Error(err)
		printer.Fatal("could not locate")
	}
}
