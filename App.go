package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	router:=mux.NewRouter()
	router.HandleFunc("/", handlerGet).Methods(http.MethodGet)
	router.HandleFunc("/", handlerPost).Methods(http.MethodPost)

	s:=http.Server{
		Addr:              "0.0.0.0:5000",
		Handler:           router,
	}
	log.Fatal(s.ListenAndServe())
}

func handlerGet(w http.ResponseWriter, r *http.Request){

	_,e:=io.WriteString(w, os.Getenv("GET_MSG"))
	if e!=nil {
		_,_=fmt.Fprint(w,e.Error())
	}

}

func handlerPost(w http.ResponseWriter, r *http.Request){
	_,e:=io.WriteString(w, os.Getenv("POST_MSG"))
	if e!=nil {
		_,_=fmt.Fprint(w,e.Error())
	}
}