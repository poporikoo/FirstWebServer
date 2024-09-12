package WebApp

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
