package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	User string `json:"user"`
	Message string `json:"message"`
	Location Location `json:"location"`
}

const(
	DISTANCE = "200km"
)

func handlerPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one post request")

	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil {
		panic(err)
	}

	fmt.Printf("Post received with message: %s\n", p.Message)
	fmt.Fprintf(w, "Post received with message: %s\n", p.Message)
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")

	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}

	fmt.Printf("range is %s", ran)
	fmt.Fprintf(w, "Search received: %f %f", lat, lon)

	p := &Post{
		User: "1111",
		Message: "woshidabendan",
		Location: Location{
			Lat: lat,
			Lon: lon,
		},

	}

	js, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func main() {
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

