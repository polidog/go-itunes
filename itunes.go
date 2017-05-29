package itunes

import (
	"encoding/json"
	"github.com/polidog/go-itunes/script"
)

type ItunesCommand func() ([]byte, error)

type Track struct {
	Album    string `json:"album"`
	Artist   string `json:"artist"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Time     string `json:"time"`
}

func (t Track) eq(check Track) bool {
	return t.Artist == check.Artist && t.Name == check.Name && t.Album == check.Album
}

type Result struct {
	Status      bool   `json:"status"`
	PlayerState string `json:"player_state"`
	Track       Track  `json:"track"`
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

func GetTrack() (Result, error) {
	return exec("track")
}

func Exec(command string) (Result, error) {
	return exec(command)
}

func exec(command string) (Result, error) {
	script, err := script.NewScript()
	result := Result{
		Track: Track{},
	}
	if err != nil {
		return result, err
	}

	str, execErr := script.Exec(command)
	if execErr != nil {
		return result, execErr
	}

	jsonErr := json.Unmarshal(str, &result)
	if jsonErr != nil {
		return result, jsonErr
	}

	return result, nil
}
