package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"

	model "./models"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	r := mux.NewRouter()

	// home page
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("template/home.html")

		checkError(err)

		t.Execute(rw, nil)
	})

	// users page
	r.HandleFunc("/users", viewUser)

	r.HandleFunc("/users/{userID}", viewUserDetail)
	// albums page
	r.HandleFunc("/albums", viewAlbum)
	// todos page
	r.HandleFunc("/todos", viewTodo)
	// posts page
	r.HandleFunc("/posts", viewPost)

	http.Handle("/", r)

	http.ListenAndServe(":9000", nil)

}

func viewUser(rw http.ResponseWriter, r *http.Request) {

	data := struct {
		Title string
		Users []model.User
	}{
		Title: "Users List",
		Users: getUsers(),
	}
	t, err := template.ParseFiles("template/users.html")

	checkError(err)

	err = t.Execute(rw, data)

	checkError(err)
}

func viewUserDetail(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Println(vars)
	userID, _ := strconv.Atoi(vars["userID"])

	user := getUser(userID)

	t, err := template.ParseFiles("template/user_detail.html")

	checkError(err)

	data := struct {
		User model.User
	}{
		User: user,
	}
	err = t.Execute(rw, data)

	checkError(err)
}

func viewAlbum(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/albums.html")

	checkError(err)

	t.Execute(rw, nil)
}

func viewTodo(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/todos.html")

	checkError(err)

	t.Execute(rw, nil)
}

func viewPost(rw http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("template/posts.html")

	checkError(err)

	t.Execute(rw, nil)
}

func getUsers() []model.User {

	var users []model.User

	res, err := http.Get(os.Getenv("API_USER"))

	checkError(err)

	defer res.Body.Close()

	output, err := ioutil.ReadAll(res.Body)

	checkError(err)

	json.Unmarshal(output, &users)

	var newUsers []model.User

	todos := getTodos()

	for _, user := range users {

		for _, todo := range todos {

			if todo.UserID == user.ID {
				user.Todos = append(user.Todos, todo)
			}
		}

		newUsers = append(newUsers, user)
	}

	return newUsers
}

func getUser(userID int) model.User {

	var newUser model.User

	users := getUsers()

	for _, user := range users {
		if user.ID == userID {
			newUser = user
		}
	}
	return newUser
}

func getTodos() []model.Todo {

	var todos []model.Todo

	res, err := http.Get(os.Getenv("API_TODO"))

	checkError(err)

	defer res.Body.Close()

	output, err := ioutil.ReadAll(res.Body)

	checkError(err)

	json.Unmarshal(output, &todos)

	return todos
}

func getUserTodos(userID int) []model.Todo {
	todos := getTodos()

	var newTodos []model.Todo

	for _, todo := range todos {
		if todo.UserID == userID {

			newTodos = append(newTodos, todo)
		}
	}

	return newTodos
}

func checkError(err error) {

	if err != nil {
		panic(err.Error())
	}
}
