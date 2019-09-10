package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SampleJSON struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	fmt.Println("in client main")
	//sj := []byte(`{"name":"ravi", "age":29}`)
	sj := SampleJSON{Name: "ravi", Age: 29}
	fmt.Println(sj)
	//var buf bytes.Buffer
	//json.NewEncoder(&buf).Encode(sj)
	requestByte, _ := json.Marshal(sj)
	resp, err := http.Post("http://localhost:8080/message", "application/json", bytes.NewBuffer(requestByte))
	//resp, err := http.Post("http://localhost:8080/message", "application/json", &buf)
	if err != nil {
		fmt.Println(err)
	}
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
