package script

import (
	"testing"
	"reflect"
)

func TestNewScript(t *testing.T) {
	script, err := NewScript()
	if err != nil {
		t.Error(err)
	}
	//fmt.Println(reflect.TypeOf(script) )
	if reflect.TypeOf(script).Name() != "AppleScript" {
		t.Error("script not created.")
	}
}

func TestNewAppleScript(t *testing.T) {
	data := make([]byte, 0)
	script, err := newAppleScript(data)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(script).Name() != "AppleScript" {
		t.Error("script not AppleScript.")
	}
}

func TestNewWindowsScript(t *testing.T) {
	data := make([]byte, 0)
	script, err := newWindowsScript(data)
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(script).Name() != "WindowsScript" {
		t.Error("script not WindowsScript.")
	}
}

func TestCreateScriptFile(t *testing.T) {
	data := make([]byte, 0)
	path, err := createScriptFile("hoge.html", data)
	if err != nil {
		t.Error(err)
	}

	if !exists(path) {
		t.Error("file not created.")
	}
}


