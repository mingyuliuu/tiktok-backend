package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Specify the max limit of the file to be 100MB
	r.ParseMultipartForm(100 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving the file: ", err)
		return
	}
	defer file.Close()

	fmt.Printf("%+v (%+v bytes) was uploaded.\n", handler.Filename, handler.Size)

	// Create a temporary file in the test-videos directory
	tempFile, err := ioutil.TempFile("test-videos", "test-upload-*.mov")
	if err != nil {
		fmt.Println("Error creating temporary local file: ", err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading uploaded file: ", err)
	}
	tempFile.Write(fileBytes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData := []byte(`{"message":"Video uploaded successfully"}`)
	w.Write(jsonData)
}
