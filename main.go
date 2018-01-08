package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

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

func main() {
	envJSON, err := json.MarshalIndent(envMap(), "", "  ")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/env.js", makeEnvHandler(envJSON))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
