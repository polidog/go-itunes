package itunes

import "time"

type Watcher struct {
	time int
	current Track
}

func (w *Watcher) Watch() bool {
	result,err := Track()
	if err == nil {
		w.current = result.Track
		return true
	}

	time.Sleep(w.time * time.Millisecond)
	if &w.current == nil || !w.current.eq(result.Track) {
		w.current = result.Track;
		return true;
	}
	return false;
}


func Watcher(time int) Watcher {
	return Watcher {
		time: time,
	}
}
