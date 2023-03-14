package main

import (
	"crypto/sha1"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
)

const (
	confFile = "/etc/fileupl/config.toml"
)

var (
	config Config
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", confFile, "Path to config file")
	flag.Parse()

	var err error
	config, err = loadConfig(configFile)
	if err != nil {
		fmt.Println("Error loading config file:", err)
		fmt.Println("Falling back to default configuration")
		fmt.Println()
	}

	if _, err := os.Stat(config.Directory); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(config.Directory, 0755)
		if err != nil {
			panic(err)
		}
	}

	if config.Password == "" {
		fmt.Println("WARNING: Password is not set.")
		fmt.Println()
	}

	fmt.Println("Files will be stored in", config.Directory)
	fmt.Println("Maximum file size is", config.MaxMB, "MB")
	fmt.Println("Service listening on", config.ListenAddress, "->", config.Url)

	// http handlers
	fs := http.FileServer(http.Dir(config.Directory))
	http.Handle("/upload", postRequirePassword(uploadHandler))
	http.Handle("/files/", http.StripPrefix("/files/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, documentation(config.Url))
	})

	panic(http.ListenAndServe(config.ListenAddress, nil))
}

func getNameAndPath(content []byte, originalName string, preserve bool) (fileName string, fullPath string) {
	if !preserve {
		fileName = fmt.Sprintf("%x", sha1.Sum(content)) + path.Ext(originalName)
	} else {
		fileName = originalName
	}
	fullPath = path.Join(config.Directory, fileName)

	return
}
