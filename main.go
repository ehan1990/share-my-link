package main

import (
	"encoding/json"
	"os"
	"log"
	"net/http"
)

type EncodeBody struct {
	Url  string `json:"url"`
	Encoded string `json:"encoded"`
}

type Response struct {
	Msg     string `json:"msg"`
	Version string `json:"version"`
}

const VERSION = "1.0.4"

func statusEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	res := Response{Msg: "server is running", Version: VERSION}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonRes)
	return
}

func randomEndpoint(w http.ResponseWriter, r *http.Request) {
	body, err := os.ReadFile("static/index.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
		log.Println("called random")
	w.Write(body)
}

func main() {
	log.Printf("running version %v", VERSION)
	http.HandleFunc("/status", statusEndpoint)
	http.HandleFunc("/random", randomEndpoint)
	http.ListenAndServe(":80", nil)
}
