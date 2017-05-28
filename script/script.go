package script

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"fmt"
)

// script
type Script interface {
	Exec(command string) ([]byte, error)
}

type ScriptFactory func(data []byte) (Script, error)

var files = map[string]string{
	"windows": "files/win/iTunes.js",
	"darwin":  "files/mac/itunes.js",
}

var factory = map[string]ScriptFactory{
	"windows": newWindowsScript,
	"darwin":  newAppleScript,
}

func NewScript() (Script, error) {
	data, err := Asset(files[runtime.GOOS])
	if err != nil {
		return nil, err
	}
	return factory[runtime.GOOS](data)
}

func newAppleScript(data []byte) (Script, error) {
	path, err := createScriptFile("mac_itunes.js", data)
	if err != nil {
		return nil, err
	}
	return AppleScript{
		path: path,
	}, nil
}

func newWindowsScript(data []byte) (Script, error) {
	path, err := createScriptFile("win_itunes.js", data)
	if err != nil {
		return nil, err
	}

	return WindowsScript{
		path: path,
	}, nil
}

func createScriptFile(name string, data []byte) (string, error) {
	dir := os.TempDir()
	fmt.Println(dir)
	path := filepath.Join(dir, name)
	if exists(path) {
		return path, nil
	}
	err := ioutil.WriteFile(path, data, 777)
	if err != nil {
		return "", err
	}
	return path, nil
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

type AppleScript struct {
	path string
}

func (a AppleScript) Exec(command string) ([]byte, error) {
	return exec.Command("osascript","-l","JavaScript", a.path, command).Output()
}

type WindowsScript struct {
	path string
}

func (a WindowsScript) Exec(command string) ([]byte, error) {
	return exec.Command("cscript -e", a.path, command).Output()
}
