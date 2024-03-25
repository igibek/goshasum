package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var (
	outputFile = flag.String("output", "", "Output file name. If missing the output will be printed to STDOUT")
	ignoreList = flag.String("ignore", ".git", "List of ignored files and directories seperated by common. Example: -ignore=.git,node_modules")
	ignores    = strings.Split(*ignoreList, ",")
	algo       = flag.String("algo", "sha256", "Hashing algorithm to use. Possible values: sha1, sha256, sha512")
)

func calculateHash(filename string) (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var hash hash.Hash

	switch *algo {
	case "sha256":
		hash = sha256.New()
	case "sha512":
		hash = sha512.New()
	case "sha1":
		hash = sha1.New()
	}

	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// Get the finalized hash result as a byte slice
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs), nil
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <path>\n", os.Args[0])
		fmt.Println("OPTIONS:")
		flag.PrintDefaults()
	}
	flag.Parse()

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
			if slices.Contains(ignores, d.Name()) {
				return filepath.SkipDir
			}

			// ignore the directories and symbolic links
			if d.IsDir() == false && d.Type()&os.ModeSymlink == 0 {
				hash, err := calculateHash(path)
				if err != nil {
					log.Fatal(err)
				}
				if *outputFile == "" {
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

		if *outputFile == "" {
			fmt.Println(hash, target)
		}
	}
}
