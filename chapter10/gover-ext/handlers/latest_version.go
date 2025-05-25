package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GoVersion struct {
	Version string `json:"version"`
}

func LatestVersionGo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Latest Go Version Called")
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://go.dev/dl/?mode=json", nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Unexpected status code: %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	// Buffer the response body for decoding
	var bodyBuffer bytes.Buffer
	tee := io.TeeReader(resp.Body, &bodyBuffer)

	// Drain the tee to ensure we can parse the body later
	if _, err := io.Copy(io.Discard, tee); err != nil {
		http.Error(w, "Failed to process response body", http.StatusInternalServerError)
		return
	}

	// Parse the JSON response
	var versions []GoVersion
	if err := json.NewDecoder(&bodyBuffer).Decode(&versions); err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		fmt.Printf("Failed to decode JSON: %v\n", err)
		return
	}

	// The latest version is the first item in the array
	if len(versions) > 0 {
		trimmedVersion := strings.TrimPrefix(versions[0].Version, "go")
		fmt.Fprintf(w, "%s", trimmedVersion) // Write the version to the HTTP response
		return
	}

	http.Error(w, "No versions found", http.StatusNotFound)
}
