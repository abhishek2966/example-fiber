package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const urlPosts = "https://jsonplaceholder.typicode.com/posts"

type post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	UsrID int    `json:"userId"`
}

func HandlePostsSave(w http.ResponseWriter, r *http.Request) {
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "post required: %v", err)
		return
	}
	var p post
	json.Unmarshal(postData, &p)
	if p.UsrID == 0 || len(p.Body) == 0 || len(p.Title) == 0 {
		http.Error(w, "Ill-formated post", http.StatusBadRequest)
		return
	}
	postData, _ = json.Marshal(p)
	body := bytes.NewReader(postData)

	req, _ := http.NewRequest(http.MethodPost, urlPosts, body)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	//io.Copy(io.Discard, resp.Body)
	responseData, _ := io.ReadAll(resp.Body)
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}
