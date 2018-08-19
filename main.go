package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

type Message struct {
	Text string
}

func main() {
	port := "8080"
	http.HandleFunc("/", handle)
	fmt.Printf("listening on port %v", port)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	msg := Message{
		"Hello Kube!!",
	}
	js, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
