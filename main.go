package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/doggy.jpg", serve)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `<h1>
				<body>
					<p>Here is a photo:</p>
					<img src="/doggy.jpg" />
				</body>
			</h1>`

	fmt.Fprintf(w, html)
}

func serve(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "doggy.jpg")
}
