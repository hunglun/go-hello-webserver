package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Set the path to the React build directory
	reactBuildDir := os.Getenv("REACT_APP_BUILD_DIR") // "./client/build"
	if reactBuildDir == "" {
		reactBuildDir = "./client/build"
	}
	// Create a file server handler for the React static files
	fs := http.FileServer(http.Dir(reactBuildDir))

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
			fmt.Println("CSS")

		}
		fmt.Println("URL: ", r.URL.Path)

		fs.ServeHTTP(w, r)
	}

	// Set the handler function for all routes
	http.HandleFunc("/", handler)

	// Get the port from the environment variable or use the default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the web server
	log.Printf("Server started on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
