package main

import "net/http"

// curl command to test
// curl -s http://localhost:8080/static/css/styles.css
func main() {
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":5001", nil)
}
