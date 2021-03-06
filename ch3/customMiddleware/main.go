package main

import (
	"fmt"
	"net/http"
)

func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing mmiddleware before request phase!")
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Executing mmiddleware after request phase!")
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Business logic here
	fmt.Println("Excuting mainHandler...")
	w.Write([]byte("OK"))
}

func main() {
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalHandler))

	http.ListenAndServe(":8000", nil)
}
