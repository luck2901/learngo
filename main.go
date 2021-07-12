package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{} //instance를 만들고

//interface사용
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Bar!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Handler 를 등록(request가 들어왔을 때 어떤 일을 할 것인지	)
		//index 페이지 경로
		fmt.Fprint(w, "Hello World")
		//write에 "hello world" 를 출력
		//response로 hello world를 준다.
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux)
	//ListenAndServe로 실행 및 구현
}
