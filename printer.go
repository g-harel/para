package main

import (
	"os"

	"github.com/fatih/color"
)

type Printer struct{}

func (p *Printer) Error(err error) {
	color.New(color.FgRed).Println(err)
}

func (p *Printer) Fatal(msg string) {
	color.New(color.Bold).Add(color.FgRed).Println(msg)
	os.Exit(1)
}
