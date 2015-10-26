package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	flag.Parse()
	baseDirectoryPath := flag.Arg(0)
	if baseDirectoryPath == "" {
		fmt.Printf("base directory was not set.\n")
		os.Exit(1)
	}
	if _, err := os.Stat(baseDirectoryPath); err != nil {
		log.Fatal(err)
	}

	fileCount := 0
	var totalFileSize int64
	totalFileSize = 0

	err := filepath.Walk(baseDirectoryPath,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				// skip hidden directory. (in prefix is dot.)
				if strings.HasPrefix(filepath.Base(path), ".") {
					return filepath.SkipDir
				}
				return nil
			}
			if strings.HasPrefix(filepath.Base(path), ".") {
				return nil
			}

			fmt.Printf("FileName: %s\n", path)
			fmt.Printf("  Size: %d\n", info.Size())
			fmt.Printf("  Modified at: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
			fmt.Printf("  Ino:%d\n", info.Sys().(*syscall.Stat_t).Ino)

			totalFileSize += info.Size()
			fileCount++

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("----\n")
	fmt.Printf("Total File Count: %d\nTotal File Size: %d\n", fileCount, totalFileSize)
	fmt.Printf("----\n")
}
