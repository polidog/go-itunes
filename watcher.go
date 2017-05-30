package itunes

import (
	"time"
)

type Watcher struct {
	time  int
	Track Track
}

func (w *Watcher) Watch() bool {
	time.Sleep(time.Duration(w.time) * time.Millisecond)
	result, err := GetTrack()

	if result.PlayerState != "playing" {
		return false
	}

	if err == nil {
		if !w.Track.eq(result.Track) {
			w.Track = result.Track
			return true
		}
	}
	return false
}

func NewWatcher(time int) Watcher {
	return Watcher{
		time:  time,
		Track: Track{},
	}
}
