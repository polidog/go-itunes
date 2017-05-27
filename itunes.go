package itunes

import (
	"encoding/json"
	"github.com/polidog/go-itunes/script"
)

type ItunesCommand func() ([]byte, error)

type Track struct {
	Album    string `json:"album"`
	Artist   string `json:"artist"`
	Category string `json:"category"`
	Time     string `json:"time"`
}

type Result struct {
	Status       bool   `json:"status"`
	PlayerStatus string `json:"player_status"`
	Track        Track  `json:"track"`
}

func Play() (Result, error) {
	return exec("play")
}

func Pause() (Result, error) {
	return exec("pause")
}

func Stop() (Result, error) {
	return exec("stop")
}

func Next() (Result, error) {
	return exec("next")
}

func Previous() (Result, error) {
	return exec("previous")
}

func Fadeout() (Result, error) {
	return exec("fadeout")
}

func Fadein() (Result, error) {
	return exec("facein")
}

func exec(command string) (Result, error) {
	script, err := script.NewScript()
	result := Result{}
	if err != nil {
		return result, err
	}

	str, execErr := script.Exec(command)
	if execErr != nil {
		return result, execErr
	}


	jsonErr := json.Unmarshal(str, result)
	if jsonErr != nil {
		return result, jsonErr
	}

	return result, nil
}
