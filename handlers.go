package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// middleware checking password for post requests
func postRequirePassword(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// we only care if a password is defined if it's POST request
		if config.Password != "" && r.Method == http.MethodPost {
			// support pw in header and form data
			ph := r.Header.Get("password")
			pf := r.FormValue("password")

			// serve request if we have a match, otherwise return error
			if ph == config.Password || pf == config.Password {
				next(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Incorrect or missing password")
			}
		} else {
			next(w, r)
		}
	})
}

// handle file uploads
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte(uploadPage))
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Client connected:", r.RemoteAddr)

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		writeError(w, "Invalid Content-Type specified", nil)
		return
	}

	err := r.ParseMultipartForm(1024 * 1024 * int64(config.MaxMB))
	if err != nil {
		writeError(w, "Error parsing multi-part form:", err)
		return
	}

	file, fHeader, err := r.FormFile("file")
	if err != nil {
		writeError(w, "Error parsing file:", err)
		return
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		writeError(w, "Error reading file:", err)
		return
	}

	preserveName := r.URL.Query().Get("pfn") != ""
	fileName, fullPath := getNameAndPath(b, fHeader.Filename, preserveName)

	err = os.WriteFile(fullPath, b, 0644)
	if err != nil {
		writeError(w, "Error storing file:", err)
		return
	}

	fmt.Println("File uploaded:", fHeader.Filename, "stored as", fileName)
	fmt.Fprintln(w, config.Url+"/files/"+fileName)
}

func writeError(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(http.StatusBadRequest)
	if err == nil {
		fmt.Fprintln(w, msg)
	} else {
		fmt.Fprintln(w, msg, err)
	}
}
