package main

import "fmt"
import "net/http"
import "io/ioutil"
import "html/template"

type camp struct {
	Title string
	Meme  string
}

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
	// api exxample: read wcnexus html file
	content := loadFile("https://www.wcnexus.com")
	fmt.Println(content)
}

// templating
func campHand(w http.ResponseWriter, r *http.Request) {
	page := camp{Title: "ChronoAmplifier", Meme: "Internal Error encountered! Time boom in 5 4 3 2 1 0 -1 -2 -3 -4 -5 ..."}

	temp, _ := template.ParseFiles("./template.html")
	temp.Execute(w, page)
}

func main() {
	http.HandleFunc("/", indexHand)
	http.HandleFunc("/gdh", gdhHand)
	http.HandleFunc("/camp", campHand)
	http.ListenAndServe(":5891", nil)
}
