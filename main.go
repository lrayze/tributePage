package main
import (
	"net/http"
	"log"
	"os"
)

/* SERVIDOR DE WEB ESTATICA */

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening http://localhost:8080...")
	http.ListenAndServe(":"+port, mux)
}
