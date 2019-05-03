package main

import "fmt"
import "net/http"
import "io/ioutil"
import "html/template"

type page struct {
	Camp    camp
	Pets    []pet
	Tickets []ticket
}

type camp struct {
	Title string
	Meme  string
}

// pet experience center
type pet struct {
	Name        string
	Type        string
	Spontaneity float32
}

type ticket struct {
	Name  string
	Class int
	Price float32
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

	thepage := page{
		Camp: camp{
			Title: "ChronoAmplifier",
			Meme:  "Internal Error encountered! Time boom in 5 4 3 2 1 0 -1 -2 -3 -4 -5 ...",
		},

		Pets: []pet{
			pet{
				Name:        "Bingganmao",
				Type:        "Cat",
				Spontaneity: 1.5,
			},
			pet{
				Name:        "JujuHu",
				Type:        "Tiger",
				Spontaneity: 7.5,
			},
			pet{
				Name:        "Honghongwang",
				Type:        "Dog",
				Spontaneity: 0.75,
			},
		},

		Tickets: []ticket{
			ticket{
				Name:  "Ultimate",
				Class: 9,
				Price: 5000,
			},
			ticket{
				Name:  "Extreme",
				Class: 8,
				Price: 4000,
			},
			ticket{
				Name:  "Immersive",
				Class: 5,
				Price: 1000,
			},
			ticket{
				Name:  "Economy",
				Class: 3,
				Price: 500,
			},
			ticket{
				Name:  "Just-look",
				Class: 1,
				Price: 100,
			},
		},
	}

	temp, _ := template.ParseFiles("./static/template.html")
	temp.Execute(w, thepage)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/api", indexHand)
	http.HandleFunc("/gdh", gdhHand)
	http.HandleFunc("/camp", campHand)
	fmt.Println("Server will be Listening on port 5891")
	http.ListenAndServe(":5891", nil)

}
