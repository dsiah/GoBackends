package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	path   string
	visits = map[string]int{}
)

func main() {
	PORT := "127.0.0.1:8080"
	http.HandleFunc("/", showCompleteTasks)
	// http.Handle("/", http.FileServer(http.Dir("/Users/dsiah/Downloads")))
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func showCompleteTasks(w http.ResponseWriter, r *http.Request) {
	path = r.URL.EscapedPath()
	count := calculateVisitorCount(visits, path)
	str := fmt.Sprintf("The path %s has been visited %d times since starting", path, count)
	log.Println(visits)
	w.Write([]byte(str))
}

func calculateVisitorCount(m map[string]int, addr string) int {
	m[addr]++
	return m[addr]
}
