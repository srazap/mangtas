package api

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"
)

func readFileBody(file *multipart.FileHeader) ([]byte, error) {

	src, err := file.Open()
	if err != nil {
		log.Println("Error2:", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(src)
	if err != nil {
		log.Println("Error3:", err)
		return nil, err
	}
	return body, nil
}

func Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error1:", err)
		return err
	}

	data, err := readFileBody(file)
	if err != nil {
		return err
	}

	words := wordCount(string(data))

	return c.JSON(http.StatusOK, words)
}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}
