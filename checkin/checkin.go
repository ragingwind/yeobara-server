package checkin

import (
	"html/template"
	"net/http"
)

type IndexInfo struct {
	Key string
}

func AppEntry(w http.ResponseWriter, r *http.Request) {
	info := IndexInfo{"testkey"}
	t, _ := template.ParseFiles("checkin/index.html")
	t.Execute(w, info)
}

func init() {
	http.HandleFunc("/", AppEntry)
	http.Handle("/static/", http.FileServer(http.Dir("./checkin")))
}
