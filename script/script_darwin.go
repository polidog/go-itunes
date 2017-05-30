package script

import (
	"os/exec"
)

func NewScript() (Script, error) {

	in := "files/mac/itunes.js"
	out := "mac_111itunes.js"

	path, err := createScriptFile(in,out)
	if err != nil {
		return nil, err
	}

	return AppleScript{
		path: path,
	}, nil
}

type AppleScript struct {
	path string
}

func (a AppleScript) Exec(command string) ([]byte, error) {
	return exec.Command("osascript", a.path, command).Output()
}
