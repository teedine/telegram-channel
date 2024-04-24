package file

import (
	"fmt"
	"math"
	"mime"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

func WatchFolder(paths []string, f func(s string)) error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer w.Close()

	go dedupLoop(w, f)

	for _, p := range paths {
		err = w.Add(p)
		if err != nil {
			return err
		}
	}

	<-make(chan struct{}) // Block forever
	return nil
}

func dedupLoop(w *fsnotify.Watcher, f func(s string)) {
	var (
		waitFor = 1 * time.Second

		// Keep track of the timers, as path → timer.
		mu     sync.Mutex
		timers = make(map[string]*time.Timer)

		// Callback we run.
		Func = func(e fsnotify.Event) {
			fmt.Printf("found event for %s\n", e.Name)

			// clean weird path that fsnotify returns
			f(strings.ReplaceAll(e.Name, "\\", "/"))

			// Don't need to remove the timer if you don't have a lot of files.
			mu.Lock()
			delete(timers, e.Name)
			mu.Unlock()
		}
	)

	for {
		select {
		case err, ok := <-w.Errors:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}
			fmt.Println(err)
		case e, ok := <-w.Events:
			if !ok { // Channel was closed (i.e. Watcher.Close() was called).
				return
			}

			// We just want to watch for file creation, so ignore everything
			// outside of Create and Write.
			if !e.Has(fsnotify.Create) && !e.Has(fsnotify.Write) {
				continue
			}

			// Get timer.
			mu.Lock()
			t, ok := timers[e.Name]
			mu.Unlock()

			// No timer yet, so create one.
			if !ok {
				t = time.AfterFunc(math.MaxInt64, func() {
					Func(e)
				})
				t.Stop()

				mu.Lock()
				timers[e.Name] = t
				mu.Unlock()
			}

			// Reset the timer for this path, so it will start from 100ms again.
			t.Reset(waitFor)
		}
	}
}

func IsDirectory(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

func cleanPath(s string) string {
	var c string

	r := []rune(s)

	for i := 0; i < len(r); i++ {
		if r[i] == '\\' {
			c = c + string(r[i])
			for k := i; k < len(r); k++ {
				if r[k] != '\\' {
					c = c + string(r[i])
					break
				} else {
					i++
				}
			}
		} else {
			c = c + string(r[i])
		}
	}

	return c
}

func GetSize(s string) int {
	f, _ := os.Stat(s)
	return int(f.Size())
}

func IsVideo(ext string) bool {
	s := mime.TypeByExtension(ext)
	if b, _, _ := strings.Cut(s, "/"); b == "video" {
		return true
	}

	return false
}
