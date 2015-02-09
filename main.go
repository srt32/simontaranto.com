package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	defaultRedirectPath = "http://blog.simontaranto.com"
)

func main() {
	http.HandleFunc("/", redirectHandler)

	port := os.Getenv("port")
	fmt.Printf("Running on port: %v", port)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Problem in ListenAndServe", err)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	var redirectedPath string
	requestedPath := r.URL.Path

	rawRedirectedPath := urlMap()[requestedPath]

	if rawRedirectedPath == "" {
		redirectedPath = defaultRedirectPath
	} else {
		redirectedPath = defaultRedirectPath + rawRedirectedPath
	}

	http.Redirect(w, r, redirectedPath, 301)
}

func urlMap() map[string]string {
	return map[string]string{
		"/bingbingbing": "/bingbingbing",
		"/foo":          "/bar",
		"/zing":         "/zong",
	}
}
