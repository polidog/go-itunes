package script

import (
	"os/exec"
	"runtime"
	"errors"
)

// script
type Script interface {
	GetFile() string
	Exec(command string) ([]byte, error)
}

func NewScript() Script {
	if runtime.GOOS == "windows" {
		return WindowsScript{
			File: "./win/iTunes.js",
		}
	} else if runtime.GOOS == "mac" {
		return AppleScript{
			File: "./mac/iTunesTransport.scpt",
		}
	}
	return NilScript{
		File: "",
	}
}


type AppleScript struct {
	File string
}

func (a AppleScript) Exec(command string) ([]byte, error) {
	return exec.Command("osascript", a.File, command).Output()
}

func (a AppleScript) GetFile() string {
	return a.File
}

type WindowsScript struct {
	File string
}

func (a WindowsScript) Exec(command string) ([]byte, error) {
	return exec.Command("cscript", a.File, command).Output()
}

func (a WindowsScript) GetFile() string {
	return a.File
}

type NilScript struct {
	File string
}

func (a NilScript) Exec(command string) ([]byte, error) {
	byte := make([]byte, 0)
	return byte, errors.New("No support os.")
}

func (a NilScript) GetFile() string {
	return a.File
}
