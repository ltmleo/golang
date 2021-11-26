package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

var host = "0.0.0.0"
var port = "8081"

func main() {
	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()                                            
	apiRouter.PathPrefix("/encode").HandlerFunc(GetEncode).Methods("GET")
	apiRouter.PathPrefix("/decode").HandlerFunc(GetDecode).Methods("GET")
	fmt.Printf("Server is running on %s:%s \n", host, port)                          
	http.ListenAndServe(":"+port, router)

}

func GetEncode(w http.ResponseWriter, r *http.Request) {
	payload := Encode(r.URL.Query().Get("string"))
	response, _ := json.Marshal(payload)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetDecode(w http.ResponseWriter, r *http.Request) {
	payload, err := Decode(r.URL.Query().Get("string"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	response, _ := json.Marshal(payload)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}