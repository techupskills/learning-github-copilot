package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/github/testdatabot/handlers"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	http.HandleFunc("/latest-version-go", handlers.LatestVersionGo)
	http.HandleFunc("/supported-versions-go", handlers.SupportedVersionsGo)
	http.HandleFunc("/is-supported-or-eol", handlers.IsSupportedOrEOL)
	http.HandleFunc("/_ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8080", nil)
	return nil
}
