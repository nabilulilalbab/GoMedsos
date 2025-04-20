package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/nabilulilalbab/library"
)




func main()  {
	router := mux.NewRouter().PathPrefix("/v1").Subrouter()
	router.HandleFunc("/image", library.CreateHandler(handleCreateImage))
  log.Println("Listen port 3001")	
	log.Fatal(http.ListenAndServe(":3001", router))	
}


func handleCreateImage(w http.ResponseWriter, r *http.Request) (int, error) {
	resp := library.NewResponse("Hello mom", nil)
  library.WriteJson(w, http.StatusCreated, resp)	
	return http.StatusCreated,  nil
}
