package script

import (
	"testing"
	"reflect"
	"runtime"
)

func TestNewScript(t *testing.T) {
	script, err := NewScript()
	if err != nil {
		t.Error(err)
	}

	name := reflect.TypeOf(script).Name() 

        switch runtime.GOOS {
	case "windows":
	    if name != "WindowsScript" {
		    t.Error("script not created.")
	    }
	case "darwin":
	    if name != "AppleScript" {
		    t.Error("script not created.")
	    }
	}
}

func TestCreateScriptFile(t *testing.T) {

	path, err := createScriptFile("empty","hoge.html")
	if err == nil {
		t.Error("this file should not exist because it does not exist")
	}



	path, err = createScriptFile("files/mac/ITunesTransport.scpt","hoge.html")
	if err != nil {
		t.Error(err)
	}

	if !exists(path) {
		t.Error("file not created.")
	}
}

