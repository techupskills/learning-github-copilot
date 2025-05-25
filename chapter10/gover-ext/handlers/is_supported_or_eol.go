package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type ReleaseInfo struct {
	ReleaseDate       string      `json:"releaseDate"`
	EOL               interface{} `json:"eol"` // Can be a string or false
	Latest            string      `json:"latest"`
	LatestReleaseDate string      `json:"latestReleaseDate"`
	LTS               bool        `json:"lts"`
}

func IsSupportedOrEOL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EOL check called")

	params := &struct {
		GoVersionOrRelease string `json:"go_version_or_release"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Regular expression to match only numbers and periods
	validPattern := regexp.MustCompile(`^\d+(\.\d+)*$`)

	var url_path string
	// Check if GoVersionOrRelease is not empty and matches the pattern
	if params.GoVersionOrRelease != "" && validPattern.MatchString(params.GoVersionOrRelease) {
		url_path = fmt.Sprintf("https://endoflife.date/api/go/%s.json", params.GoVersionOrRelease)
	} else {
		http.Error(w, "Invalid version specified", http.StatusNotFound)
	}

	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, url_path, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Fprintf(w, "No versions found matching %s.\n", params.GoVersionOrRelease)
		return
	}

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

	// Create a Decoder
	decoder := json.NewDecoder(strings.NewReader(bodyBuffer.String()))

	// Decode JSON into the struct
	var releaseInfo ReleaseInfo
	if err := decoder.Decode(&releaseInfo); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Handle the "eol" field based on its type
	if eolDate, ok := releaseInfo.EOL.(string); ok {
		fmt.Fprintf(w, "Version %s is EOL as of %s. The latest release was %s on %s.\n",
			params.GoVersionOrRelease, eolDate, releaseInfo.Latest, releaseInfo.LatestReleaseDate)
		return
	} else if eolBool, ok := releaseInfo.EOL.(bool); ok && !eolBool {
		fmt.Fprintf(w, "Version %s is not EOL.\n", params.GoVersionOrRelease)
		return
	} else {
		fmt.Println("EOL: unknown type")
	}

	http.Error(w, "An unexpected error occurred", http.StatusInternalServerError)
}
