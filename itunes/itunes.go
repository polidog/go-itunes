package main

import (
	"github.com/polidog/go-itunes"
	"fmt"
	"os"
)

func main() {
	var command string
	if len(os.Args) > 0 {
		command = os.Args[1]
	}

	res, err := itunes.Exec(command)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("status:", res.Status)
	fmt.Println("player_state:", res.PlayerState)

	fmt.Println("----------track-----------")
	fmt.Println("Album:", res.Track.Album)
	fmt.Println("Name:", res.Track.Name)
	fmt.Println("Artist:", res.Track.Artist)
	fmt.Println("Category:", res.Track.Category)
	fmt.Println("Time:", res.Track.Time)
}
