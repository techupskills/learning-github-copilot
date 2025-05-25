package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type VersionInfo struct {
	Cycle       string      `json:"cycle"`
	ReleaseDate string      `json:"releaseDate"`
	EOL         interface{} `json:"eol"`
}

func SupportedVersionsGo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Supported Go Versions Called")
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://endoflife.date/api/go.json", nil)
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

	// Parse the JSON into a slice of VersionInfo
	var versions []VersionInfo
	if err := json.NewDecoder(&bodyBuffer).Decode(&versions); err != nil {
		http.Error(w, "Failed to parse JSON response", http.StatusInternalServerError)
		fmt.Printf("Failed to decode JSON: %v\n", err)
		return
	}

	// Filter versions with eol == false and collect the results
	result := make([]struct {
		Cycle       string `json:"cycle"`
		ReleaseDate string `json:"releaseDate"`
	}, 0)

	for _, v := range versions {
		// Check if EOL is a boolean and is false
		if eolBool, ok := v.EOL.(bool); ok && !eolBool {
			result = append(result, struct {
				Cycle       string `json:"cycle"`
				ReleaseDate string `json:"releaseDate"`
			}{
				Cycle:       v.Cycle,
				ReleaseDate: v.ReleaseDate,
			})
		}
	}

	// Convert the result to a single string
	var builder strings.Builder
	for _, r := range result {
		builder.WriteString(fmt.Sprintf("Version: %s, Release Date: %s\n", r.Cycle, r.ReleaseDate))
	}

	// The latest version is the first item in the array
	if len(builder.String()) > 0 {
		fmt.Fprintf(w, "%s", builder.String()) // Write the list of supported versions to the HTTP response
		return
	}

	http.Error(w, "No versions found", http.StatusNotFound)
}
