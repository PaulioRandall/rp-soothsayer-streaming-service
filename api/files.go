package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const dir string = "./testdata"

var supportedVideoFormats = [...]string{
	".mp4",
}

func attachHttpFileListHandler() {
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		appendCorsHeaders(w)

		if r.Method == "OPTIONS" {
			return
		}

		files, e := listVideoFilesInDir(dir)
		if e != nil {
			http.Error(w, "Unable to access file list", http.StatusInternalServerError)
			return
		}

		json, e := json.Marshal(files)
		if e != nil {
			http.Error(w, "Unable to jsonify file list", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Content-Length", strconv.Itoa(len(json)))
		w.Write(json)
	})
}

func attachHttpFileHandler() {
	fs := http.FileServer(http.Dir(dir))

	http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		appendCorsHeaders(w)

		if r.Method == "OPTIONS" {
			return
		}

		_, r.URL.Path = shiftPath(r.URL.Path)
		fs.ServeHTTP(w, r)
	})
}

// shiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
//
// Source: Merovius (Axel Wagner) - https://blog.merovius.de/posts/2017-06-18-how-not-to-use-an-http-router/
func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

// listVideoFilesInDir reads a directory and returns the names
// of video files as a slice of strings.
func listVideoFilesInDir(dir string) ([]string, error) {
	files, e := os.ReadDir(dir)

	if e != nil {
		return nil, e
	}

	result := []string{}

	for _, f := range files {
		name := f.Name()
		if isSupportedVideoFormat(name) {
			result = append(result, name)
		}
	}

	return result, nil
}

func isSupportedVideoFormat(filename string) bool {
	ext := filepath.Ext(filename)

	for _, vidFmt := range supportedVideoFormats {
		if ext == vidFmt {
			return true
		}
	}

	return false
}
