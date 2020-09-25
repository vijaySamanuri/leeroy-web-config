package main

import (
	"io"
	"net/http"
        "os"
        "errors"

	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
        leeroyAppHostName := os.Getenv("LEEROY_APP_HOSTNAME")
        if len(leeroyAppHostName) == 0 {
                panic(errors.New("LEEROY_APP_HOSTNAME is not set"))
        }
	resp, err := http.Get("http://" + leeroyAppHostName + ":50051")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		panic(err)
	}
}

func main() {
	log.Print("leeroy web server ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
