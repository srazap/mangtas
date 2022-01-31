package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	url  = "http://localhost:1323/upload"
	name = "Golang_Test.txt"
)

func main() {

	// read from file and create form data to send in api request
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()
		part, err := m.CreateFormFile("file", name)
		if err != nil {
			return
		}
		file, err := os.Open(name)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()

	// call word count api
	resp, err := http.Post(url, m.FormDataContentType(), r)
	if err != nil {
		panic(err)
	}

	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshal into string map
	wordMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &wordMap); err != nil {
		panic(err)
	}

	// print response from api
	op, err := json.MarshalIndent(wordMap, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("===============================================================")
	fmt.Println(string(op))
	fmt.Println("===============================================================")
}
