package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	hClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//return nil
			return http.ErrUseLastResponse
		},
	}

	file, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", "upload.png")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatal(err)
	}

	contentType := writer.FormDataContentType()

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:4567/file_upload", body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", contentType)

	res, err := hClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
}
