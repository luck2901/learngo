package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{} //instance를 만들고

//interface사용
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) // Json형태로 파싱
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user) //user를 json형태로 다시 바꿈

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))

}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
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
	//위의 HandleFunc는 Handler를 Function으로 지정할 때 사용
	//Handle은 Handler라는 Instance 형태로 지정할 때는 Handle 사용

	http.ListenAndServe(":3000", mux)
	//ListenAndServe로 실행 및 구현
}
