package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	err = watcher.Add("html")
	if err != nil {
		panic(err)
	}

	fmt.Println("Watching for changes in html/...")

	var lastRun time.Time

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Rename) != 0 {
				if time.Since(lastRun) < 1*time.Second {
					continue
				}
				lastRun = time.Now()

				fmt.Println("Change detected:", event.Name)
				runBuilder()
			}
		case err := <-watcher.Errors:
			fmt.Println("Watcher error:", err)
		}
	}
}

func runBuilder() {
	cmd := exec.Command("go", "run", "builder/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running builder:", err)
	}
}
