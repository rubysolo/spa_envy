package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// return process environment formatted as javascript object
func envMap() map[string]string {
	env := make(map[string]string)
	for _, pair := range os.Environ() {
		parts := strings.SplitN(pair, "=", 2)
		key := parts[0]
		val := parts[1]
		env[key] = val
	}
	return env
}

func makeEnvHandler(envJSON []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "var config = ")
		w.Write(envJSON)
		io.WriteString(w, ";")
	}
}

// wrapped http.FileServer that can serve index.html for any
// 404 (to support SPA deep-links)
type StaticHandler struct {
	fileServer http.Handler
}

type response struct {
	buffer bytes.Buffer
	status int
	header http.Header
}

func (w *response) Write(p []byte) (int, error) {
	return w.buffer.Write(p)
}
func (w *response) WriteHeader(status int) {
	w.status = status
}
func (w *response) Header() http.Header {
	return w.header
}

func (h *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wrapped := &response{header: w.Header()}

	// pass request to native file server
	h.fileServer.ServeHTTP(wrapped, r)

	// intercept 404
	if wrapped.status == http.StatusNotFound {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
		http.ServeFile(w, r, "static/index.html")
		return
	}

	w.WriteHeader(wrapped.status)
	io.Copy(w, &wrapped.buffer)
}

func main() {
	envJSON, err := json.MarshalIndent(envMap(), "", "  ")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/env.js", makeEnvHandler(envJSON))

	fs := &StaticHandler{fileServer: http.FileServer(http.Dir("static"))}
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
