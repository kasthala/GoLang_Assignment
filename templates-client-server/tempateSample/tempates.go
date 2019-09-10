package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//PageData is pagedata
type pageData struct {
	PageTitle string
	Message   string
}

func main() {
	fmt.Println("start template")
	tm, _ := template.ParseGlob("./tempates/*.html")
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		p := []pageData{}
		for i := 1; i < 10; i++ {
			p = append(p, pageData{"golang" + strconv.Itoa(i), "lets learn it : " + strconv.Itoa(i)})
		}
		tm.ExecuteTemplate(w, "sample1.html", p)
	})

	http.HandleFunc("/bar1", func(w http.ResponseWriter, r *http.Request) {
		p := []pageData{}
		for i := 1; i < 20; i++ {
			p = append(p, pageData{"golang" + strconv.Itoa(i), "lets learn it : " + strconv.Itoa(i)})
		}
		tm.Execute(w, p)
	})
	http.ListenAndServe(":8080", nil)
}
