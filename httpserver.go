package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiMessage struct {
	Message string `json:message`
}

func getMessage(rw http.ResponseWriter, r *http.Request) {
	arr := mux.Vars(r)

	fmt.Println(arr)

	message := ApiMessage{"Istenilen id:" + arr["id"]}

	output, err := json.Marshal(message)

	if err != nil {
		fmt.Fprintf(rw, "Bir sorun oluştu!!")
	}

	fmt.Fprintf(rw, string(output))
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", getMessage)

	http.Handle("/", r)
	http.ListenAndServe(":9000", nil)
}
