package checkin

import (
	"html/template"
	"net/http"
	"encoding/json"
	"os"
	"log"
)

type IndexInfo struct {
	Meetup string
	UserId string
	Location string
}

type Config struct {
	FBAppName string
}

var config Config

func loadConfiguration() {
	file, _ := os.Open("env.json")
	err := json.NewDecoder(file).Decode(&config)

	if err != nil {
		log.Println("error:", err)
	}
}

func appEntry(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	t := template.Must(template.New("index.html").Delims("[[", "]]").ParseFiles("checkin/index.html"))
	t.Execute(w, IndexInfo{r.Form.Get("meetup"), r.Form.Get("user-id"), config.FBAppName})
}

func init() {
	loadConfiguration()

	http.HandleFunc("/", appEntry)
}
