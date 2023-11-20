package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const urlPhotos = "https://jsonplaceholder.typicode.com/photos"

type photo struct {
	AlbumID    int    `json:"albumId"`
	PhotoID    int    `json:"id"`
	PhotoTitle string `json:"title"`
	PhotoURL   string `json:"url"`
}

func HandlePhotosFetch(w http.ResponseWriter, r *http.Request) {
	albumIDString := r.URL.Query().Get("album_id")
	albumID, err := strconv.Atoi(albumIDString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "album_id required: %v", err)
		return
	}

	resp, err := client.Get(urlPhotos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "couldn't read response body: %v", err)
		return
	}

	pics := []photo{}
	json.Unmarshal(data, &pics)

	picsFiltered := []photo{}
	for _, pic := range pics {
		if pic.AlbumID == albumID {
			picsFiltered = append(picsFiltered, pic)
		}
	}
	responseData, err := json.Marshal(picsFiltered)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Marshal error: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
