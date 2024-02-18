package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		he, _ := strconv.Atoi(r.URL.Query().Get("height"))
		we, _ := strconv.Atoi(r.URL.Query().Get("weight"))
		bmi := float64(we) / ((float64(he) / 100) * (float64(he) / 100))
		fmt.Fprintf(w, "你的身高為%d，體重為%d，BMI值為%f", he, we, bmi)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
