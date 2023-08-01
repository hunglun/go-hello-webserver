package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Data struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func setCustomPath() {
	// Specify your custom PATH here
	customPath := "/usr/local/bin:/usr/bin:/bin"

	// Get the current environment PATH
	currentPath := os.Getenv("PATH")

	// Append the custom path to the current path, separating them with a colon (for Unix-like systems)
	newPath := currentPath + ":" + customPath

	// Set the new environment PATH
	os.Setenv("PATH", newPath)
}

func getSystemTime() string {
	cmd := exec.Command("date")
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}

	uptimeInfo := strings.TrimSpace(string(output))
	return uptimeInfo
}

func runCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}

	return strings.TrimSpace(string(output))
}

func main() {

	setCustomPath()

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

	// set the handler for system info
	http.HandleFunc("/deviceSystemInfo", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Entries: []Entry{
				{Name: "device name", Value: runCommand("uname", "-a")},
				{Name: "system time", Value: runCommand("date")},
				{Name: "uptime", Value: runCommand("uptime")},
				{Name: "who", Value: runCommand("w", "-h")},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	// Get the port from the environment variable or use the default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the web server
	log.Printf("Server started on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
