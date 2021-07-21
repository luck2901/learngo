package main

import (
	"net/http"

	"github.com/luck2901/learngo/Part3/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
