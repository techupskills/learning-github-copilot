package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User Called")
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://randomuser.me/api", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		fmt.Println("Something unrecoverable went wrong")
		return
	}
}
