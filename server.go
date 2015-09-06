package main

import (
	//"io/ioutil"
	"log"
	"net/http"
)

func main() {
	context := Context{"/tmp/restic"}

	repo, _ := context.Repository("user")
	repo.Init()

	errc := context.Init()
	if errc != nil {
		log.Println("context initialization failed")
		return
	}

	router := Router{context}
	port := ":8000"
	log.Printf("start server on port %s", port)
	http.ListenAndServe(port, router)
}
