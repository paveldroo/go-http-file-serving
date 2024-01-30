package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("."))))
}

// Here is an old manual fileserver

//func main() {
//	http.HandleFunc("/", index)
//	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
//	log.Fatal(http.ListenAndServe(":8000", nil))
//}
//
//func index(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	html := `<h1>
//				<body>
//					<p>Here is a photo:</p>
//					<img src="/resources/doggy.jpg" />
//				</body>
//			</h1>`
//
//	io.WriteString(w, html)
//}
//
//func serve(w http.ResponseWriter, r *http.Request) {
//	http.ServeFile(w, r, "doggy.jpg")
//}
