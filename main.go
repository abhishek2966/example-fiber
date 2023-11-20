package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhishek2966/example/config"
	"github.com/abhishek2966/example/handler"
	"github.com/gorilla/mux"
)

func main() {
	data := config.Data{}
	data.InitFlag()
	yamldata, err := data.ReadYAML()
	if err != nil {
		panic(err)
	}
	err = data.DecodeYAML(yamldata)
	if err != nil {
		log.Print(err)
	}
	addr := fmt.Sprintf(":%v", data.Port)

	r := mux.NewRouter()
	r.HandleFunc("/photos", handler.HandlePhotosFetch).Methods("GET")
	r.HandleFunc("/posts", handler.HandlePostsSave).Methods("POST")
	http.ListenAndServe(addr, r)
	fmt.Println(data)
}
