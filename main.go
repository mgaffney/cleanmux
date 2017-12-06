package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(u.HomeDir, ".tmux/resurrect")

	var deleted int
	t := time.Now().AddDate(0, 0, -14)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path == dir {
			return nil
		}
		if info.IsDir() {
			return filepath.SkipDir
		}
		if info.ModTime().Before(t) && info.Mode()&os.ModeSymlink == 0 {
			if err := os.Remove(path); err != nil {
				return err
			} else {
				deleted++
			}
		}
		return nil
	})
	fmt.Printf("%d files removed\n", deleted)
}
