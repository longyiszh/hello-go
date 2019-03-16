package main

import "fmt"
import "net/http"
import "io/ioutil"

func loadFile(path string) string {
	resp, _ := http.Get(path)
	bytes, _ := ioutil.ReadAll(resp.Body)
	content := string(bytes)
	resp.Body.Close()
	return content
}

func indexHand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go server works!")
}

func gdhHand(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "gdh works!")

	// api exxample: read wcnexus html file
	content := loadFile("https://www.wcnexus.com")
	fmt.Println(content)

}

func mainsv() {
	http.HandleFunc("/", indexHand)
	http.HandleFunc("/gdh", gdhHand)
	http.ListenAndServe(":5891", nil)
}
