package script

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func NewScript() (Script, error) {

	in := "files/mac/ITunesTransport.scpt"
	out := "mac_itunes.scpt"

	path, err := createScriptFile(in,out)
	if err != nil {
		return nil, err
	}

	return WindowsScript{
		path: path,
	}, nil
}

type AppleScript struct {
	path string
}

func (a AppleScript) Exec(command string) ([]byte, error) {
	return exec.Command("osascript", a.path, command).Output()
}
