package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/query/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		//這邊就把 name 帶入 UserInfo
		u := &UserInfo{
			Name: vars["name"],
			Age:  18,
		}
		b, err := json.Marshal(u)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	})

	//聽8000 port
	log.Fatal(http.ListenAndServe(":8000", r))
}
