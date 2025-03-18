// Recursive File Watcher
// This program implements a file system watcher that recursively monitors
// a directory and all its subdirectories for changes. It also automatically
// starts watching any newly created directories.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	// fsnotify provides native file system notifications for Go
	"github.com/fsnotify/fsnotify"
)

// watchDir recursively adds directories to the file watcher
// Parameters:
// - path: the root directory path to start watching from
// - watcher: the fsnotify.Watcher instance to use for watching
// Returns an error if any directory cannot be added to the watch list
func watchDir(path string, watcher *fsnotify.Watcher) error {
	// Use filepath.Walk to recursively traverse the directory tree
	return filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		// If there's an error accessing the path, return it
		if err != nil {
			return err
		}
		// If the current path is a directory, add it to the watch list
		if fi.Mode().IsDir() {
			fmt.Printf("Watching directory: %s\n", path)
			return watcher.Add(path)
		}
		// Skip files (we only watch directories)
		return nil
	})
}

func main() {
	// Create a new file system watcher instance
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("ERROR creating watcher: %v\n", err)
		return
	}
	// Ensure the watcher is closed when the program exits
	defer watcher.Close()

	// Start watching the initial directory structure recursively
	if err := watchDir("./", watcher); err != nil {
		fmt.Printf("ERROR watching directory: %v\n", err)
		return
	}

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
				fmt.Printf("EVENT! %s: %s\n", event.Op, event.Name)

				// Special handling for newly created directories
				// If a new directory is created, we need to start watching it too
				if event.Op&fsnotify.Create == fsnotify.Create {
					// Get information about the created file/directory
					fi, err := os.Stat(event.Name)
					if err == nil && fi.Mode().IsDir() {
						fmt.Printf("New directory created: %s\n", event.Name)
						// Add the new directory to our watch list
						if err := watchDir(event.Name, watcher); err != nil {
							fmt.Printf("ERROR watching new directory: %v\n", err)
						}
					}
				}

			// Handle any errors that occur during watching
			case err := <-watcher.Errors:
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	}()

	fmt.Println("File watcher started. Press Ctrl+C to stop...")
	// Block forever until program is terminated
	<-done
}
