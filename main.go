package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func calculateHash(filename string) (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()

	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// Get the finalized hash result as a byte slice
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

func main() {
	var outputFlag string
	var ignoreFlag string

	flag.StringVar(&outputFlag, "output", "", "Output file name. If missing it will output SHA256SUM to stdout")
	flag.StringVar(&ignoreFlag, "ignore", ".git", "List of ignored directories and/or files separated by comma (,). \nFor example: .git,node_modules")

	flag.Parse()

	ignored := strings.Split(ignoreFlag, ",")

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		log.Fatal("[PATH] is required")
	}

	target := os.Args[1]

	fileInfo, err := os.Stat(target)

	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.IsDir() {
		filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
			if slices.Contains(ignored, d.Name()) {
				return filepath.SkipDir
			}

			// ignore the directories and symbolic links
			if d.IsDir() == false && d.Type()&os.ModeSymlink == 0 {
				hash, err := calculateHash(path)
				if err != nil {
					log.Fatal(err)
				}
				if outputFlag == "" {
					fmt.Println(hash, path)
				}
			}

			return nil
		})
	} else {
		hash, err := calculateHash(target)
		if err != nil {
			log.Fatal(err)
		}

		if outputFlag == "" {
			fmt.Println(hash, target)
		}
	}
}
