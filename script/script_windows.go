package script

import (
	"os/exec"
)

func NewScript() (Script, error) {

        in := "files/win/iTunes.js"
	out := "win_itunes.js"

	path, err := createScriptFile(in,out)
	if err != nil {
		return nil, err
	}

	return WindowsScript{
		path: path,
	}, nil
}

type WindowsScript struct {
	path string
}

func (a WindowsScript) Exec(command string) ([]byte, error) {
	return exec.Command("cscript -e", a.path, command).Output()
}
