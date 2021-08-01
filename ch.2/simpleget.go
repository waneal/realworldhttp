package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	// values := url.Values{
	// 	"query": {"hello world"},
	// }

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Hitoshi Fujita")

	part := make(textproto.MIMEHeader)
	part.Set("Context-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="thumbnail.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	// fileWriter, err := writer.CreateFormFile("thumbnail","thumbnail.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("thumbnail.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
}
