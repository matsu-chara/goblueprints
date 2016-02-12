package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/matsu-chara/goblueprints/chapter7/meander"
)

func main() {
	key := os.Getenv("PLACE_API_KEY")
	if key == "" {
		log.Fatalln("PLACE_API_KEYが環境変数から取得できませんでした")
		return
	}
	meander.APIKey = key
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
