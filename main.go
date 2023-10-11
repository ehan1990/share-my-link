package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
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

func encodeEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(req.Body)
    var body EncodeBody
    err := decoder.Decode(&body)

	if err != nil {
        panic(err)
    }

	b64Encoded := b64.URLEncoding.EncodeToString([]byte(body.Url))
    body.Encoded = fmt.Sprintf("http://35.91.228.41/%s", b64Encoded)

	res, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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

func redirectToURL(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path[1:])
	newsLink, err := b64.URLEncoding.DecodeString(r.URL.Path[1:])

	if err != nil {
		panic(err)
	}
	
	log.Printf("%v\n", string(newsLink))
	http.Redirect(w, r, string(newsLink), http.StatusMovedPermanently)
}

func main() {
	log.Printf("running version %v", VERSION)
	http.HandleFunc("/status", statusEndpoint)
	http.HandleFunc("/encode", encodeEndpoint)
	http.HandleFunc("/random", randomEndpoint)
	// http.HandleFunc("/", redirectToURL)
	http.ListenAndServe(":80", nil)
}
