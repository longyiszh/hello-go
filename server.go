package main

import "fmt"
import "net/http"

func indexHand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go server works!")
}

func gdhHand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "gdh works!")
}

func main2() {
	http.HandleFunc("/", indexHand)
	http.HandleFunc("/gdh", gdhHand)
	http.ListenAndServe(":5891", nil)
}
