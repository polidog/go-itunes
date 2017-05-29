package main

import (
	"fmt"
	"github.com/polidog/go-itunes"
)

func main() {
	watcher := itunes.NewWatcher(800)
	for {
		if watcher.Watch() {
			fmt.Println("change track")
			fmt.Println(watcher.Track)
		}
	}
}
