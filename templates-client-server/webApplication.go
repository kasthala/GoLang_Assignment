package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

//SampleJSON xvsdf
type SampleJSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am called")
	//x, _ := ioutil.ReadAll(r.Body)
	t := &SampleJSON{}
	log.Println(t)
	//json.Unmarshal(x, t)
	json.NewDecoder(r.Body).Decode(t)
	log.Println(t)
	//fmt.Println(b)
	fmt.Println(t.Age)
	w.Write([]byte("hi ravi\n"))
}
func main() {

	fmt.Println("in main")

	http.HandleFunc("/message", messageHandler)

	fs := http.FileServer(http.Dir("./template"))

	http.Handle("/t/", http.StripPrefix("/t/", fs))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":8080", nil)
}
