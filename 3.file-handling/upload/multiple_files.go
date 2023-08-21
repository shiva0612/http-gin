package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

func saveForm2(c *gin.Context) {
	// Retrieve the values of the form
	_ = c.PostForm("name")
	_ = c.PostForm("age")
	_ = c.PostForm("city")

	// Retrieve the files uploaded in the form
	form, err := c.MultipartForm()
	if err != nil {
		// Handle error
		return
	}

	for k, v := range form.Value {
		fmt.Println("form key values can be fetched like this also", k, v)
	}

	for field_name, fileHeader := range form.File {

		_ = field_name
		//considering that each field has only 1 file
		file, err := fileHeader[0].Open()
		if err != nil {
			// Handle error
			return
		}
		defer file.Close()

		// Save the uploaded file
		err = saveFile2(file, fileHeader[0].Filename)
		if err != nil {
			// Handle error
			return
		}
	}

	// Do something with the form data and files
	// ...
}

func saveFile2(file multipart.File, filename string) error {
	// Open a new file with the specified filename
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Copy the contents of the uploaded file to the new file
	_, err = io.Copy(outFile, file)
	if err != nil {
		return err
	}

	return nil
}
