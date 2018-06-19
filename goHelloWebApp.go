package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"

	"github.com/ajays20078/go-http-logger"
	"github.com/gorilla/mux"

)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func index(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, "For as long as he could remember, he’d suffered from a vague nagging feeling of being not all there.")
}

func main() {
	listenPort := "3000"
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")

	fmt.Printf("Listening on :" + listenPort + "\n")
	if err := http.ListenAndServe(":"+listenPort,
		httpLogger.WriteLog(r, os.Stdout)); err != nil {
		log.Fatal(err)
	}
}
