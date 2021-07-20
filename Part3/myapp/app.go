package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func getUserInfo89Handler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r) //request 넣어주면 알아서 id 파싱
	fmt.Fprint(w, "User Id:", vars["id"])
}

//NewHandler make a new myapp handler
func NewHandler() http.Handler {
	mux := mux.NewRouter() //대용
	//mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfo89Handler)

	return mux
}
