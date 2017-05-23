package itunes

import "github.com/polidog/go-itunes/script"

type ItunesCommand func() ([]byte, error)

func Play() ([]byte, error) {
	return exec("play")
}

func Pause() ([]byte, error) {
	return exec("pause")
}

func Stop() ([]byte, error) {
	return exec("stop")
}

func Next() ([]byte, error) {
	return exec("next")
}

func Previous() ([]byte, error) {
	return exec("previous")
}

func Fadeout() ([]byte, error) {
	return exec("fadeout")
}

func Fadein() ([]byte, error) {
	return exec("facein")
}

func exec(command string) ([]byte, error) {
	script, err := script.NewScript()
	if err != nil {
		return nil, err
	}
	return script.Exec(command)
}
