package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	Message string `json:message`
}

type User struct {
	Name    string `json:name`
	Surname string `json:surname`
	Age     int    `json:age`
}

func main() {

	rootPath := "/api"
	//.../api
	http.HandleFunc(rootPath, func(rw http.ResponseWriter, r *http.Request) {

		api := API{"Merhaba"}

		output, err := json.Marshal(api)

		checkError(err)

		// rw.Header().Set("Content-Type", "aplication/json")

		fmt.Fprintf(rw, string(output))

	})

	//.../api/users
	http.HandleFunc(rootPath+"/users", func(rw http.ResponseWriter, r *http.Request) {

		users := []User{
			User{Name: "Muhammet", Surname: "ÖZATA", Age: 10},
			User{Name: "Meryam", Surname: "ÖZATA", Age: 20},
			User{Name: "Hakan", Surname: "Test", Age: 30},
		}

		output, err := json.Marshal(users)

		checkError(err)

		// rw.Header().Set("Content-Type", "aplication/json")

		fmt.Fprintf(rw, string(output))

	})

	http.HandleFunc(rootPath+"/me", func(rw http.ResponseWriter, r *http.Request) {
		user := User{Name: "Muhammet", Surname: "ÖZATA", Age: 30}

		output, err := json.Marshal(user)

		checkError(err)

		fmt.Fprintf(rw, string(output))
	})

	err := http.ListenAndServe(":9000", nil)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
