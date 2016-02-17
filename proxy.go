package main // import "github.com/gwwfps/simplerp"

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	b := os.Getenv("BACKEND_URL")

	if b == "" {
		log.Fatalln("BACKEND_URL environment variable is not set.")
	}

	u, err := url.ParseRequestURI(b)

	if err != nil {
		log.Fatalln("BACKEND_URL environment variable is not a valid URL.", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	log.Printf("Proxying %s on 8080...\n", b)
	http.ListenAndServe(":8080", proxy)
}
