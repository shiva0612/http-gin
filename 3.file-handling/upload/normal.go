package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

func saveForm(c *gin.Context) {
	// Retrieve the values of the form
	_ = c.PostForm("name")
	_ = c.PostForm("age")
	_ = c.PostForm("city")

	// Retrieve the files uploaded in the form
	aadharHeader, err := c.FormFile("aadhar")
	if err != nil {
		// Handle error
		return
	}

	panHeader, err := c.FormFile("pan")
	if err != nil {
		// Handle error
		return
	}

	saveFile(aadharHeader)
	saveFile(panHeader)
}

func saveFile(header *multipart.FileHeader) error {
	// Open a new file with the specified filename
	outFile, err := os.Create(header.Filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Copy the contents of the uploaded file to the new file
	file, err := header.Open()
	if err != nil {
		return fmt.Errorf("cannot read the file %s", header.Filename)
	}
	defer file.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		return err
	}

	return nil
}
