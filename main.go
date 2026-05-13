package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	file_server := http.FileServer(http.Dir("."))
	mux.Handle("/", file_server)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
