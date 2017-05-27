package main

import (
	"flag"
	"github.com/polidog/go-itunes"
	"fmt"
)

func main() {
	var command string
	flag.StringVar(&command, "f", "", "config file name.")
	flag.Parse();
	res, err := itunes.Exec(command)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("status:", res.Status)
	fmt.Println("player_state:", res.PlayerState)
	if res.Track {
		fmt.Println("----------track-----------")
		fmt.Println("Album:", res.Track.Album)
		fmt.Println("Artist:", res.Track.Artist)
		fmt.Println("Category:", res.Track.Category)
		fmt.Println("Time:", res.Track.Time)
	}
}
