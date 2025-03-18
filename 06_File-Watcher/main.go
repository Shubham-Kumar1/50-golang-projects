// Basic File Watcher
// This program demonstrates a simple file system watcher that monitors a single directory
// for any changes (create, modify, delete, rename events).
package main

import (
	"fmt"

	// fsnotify provides native file system notifications for Go
	"github.com/fsnotify/fsnotify"
)

func main() {
	// Create a new file system watcher instance
	// This watcher will be responsible for monitoring file system events
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
	}
	// Ensure we close the watcher when the program exits
	defer watcher.Close()

	// Create a channel to keep the program running
	// This channel will never receive a value, effectively making the program run until interrupted
	done := make(chan bool)

	// Start a goroutine to handle file system events
	go func() {
		// Infinite loop to continuously watch for events
		for {
			// Use select to handle multiple channel operations
			select {
			// Handle file system events (create, write, remove, rename, chmod)
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)
				// event.Op will be one of:
				// - Create (file/directory created)
				// - Write (file/directory modified)
				// - Remove (file/directory deleted)
				// - Rename (file/directory renamed)
				// - Chmod (file/directory permissions changed)

			// Handle any errors that occur during watching
			case err := <-watcher.Errors:
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	}()

	// Add the current directory ("./" means current directory) to the watch list
	// This tells the watcher to monitor this directory for changes
	if err := watcher.Add("./"); err != nil {
		fmt.Println("ERROR adding directory to watch:", err)
	}

	// Block forever until program is terminated
	<-done
}
