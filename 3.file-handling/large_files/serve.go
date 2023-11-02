package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func server_bufio() {
	// initialize a new Gin router
	r := gin.Default()

	// define a GET endpoint that serves a file
	r.GET("/movies/:filename", func(c *gin.Context) {
		// retrieve the filename from the URL parameter
		filename := c.Param("filename")

		// open the file on the server
		file, err := os.Open(filename)
		if err != nil {
			log.Println(err)
			c.JSON(404, gin.H{"message": "File not found"})
			return
		}
		defer file.Close()

		// get the file's size
		fileInfo, err := file.Stat()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		fileSize := fileInfo.Size()

		// set the Content-Disposition header to force a download
		c.Header("Content-Disposition", "attachment; filename="+filename)

		// set the Content-Type header based on the file extension
		// (you may need to add more types here depending on your use case)
		switch {
		case hasSuffix(filename, ".mp4"):
			c.Header("Content-Type", "video/mp4")
		case hasSuffix(filename, ".jpg") || hasSuffix(filename, ".jpeg"):
			c.Header("Content-Type", "image/jpeg")
		case hasSuffix(filename, ".png"):
			c.Header("Content-Type", "image/png")
		default:
			c.Header("Content-Type", "application/octet-stream")
		}

		// set the Content-Length header
		c.Header("Content-Length", strconv.FormatInt(fileSize, 10))

		// create a buffered reader with a buffer size of 1MB
		bufferSize := 1024 * 1024 * 8 // 1MB buffer size
		bufferedReader := bufio.NewReaderSize(file, bufferSize)

		// write the file contents to the response using the buffered reader's WriteTo method
		_, err = bufferedReader.WriteTo(c.Writer)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
	})

	// start the server
	r.Run(":8080")
}

func serve_loop() {
	// initialize a new Gin router
	r := gin.Default()

	// define a GET endpoint that serves a file
	r.GET("/movies/:filename", func(c *gin.Context) {
		// retrieve the filename from the URL parameter
		filename := c.Param("filename")

		// open the file on the server
		file, err := os.Open(filename)
		if err != nil {
			log.Println(err)
			c.JSON(404, gin.H{"message": "File not found"})
			return
		}
		defer file.Close()

		// get the file's size
		fileInfo, err := file.Stat()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		fileSize := fileInfo.Size()

		// set the Content-Disposition header to force a download
		c.Header("Content-Disposition", "attachment; filename="+filename)

		// set the Content-Type header based on the file extension
		// (you may need to add more types here depending on your use case)
		switch {
		case hasSuffix(filename, ".mp4"):
			c.Header("Content-Type", "video/mp4")
		case hasSuffix(filename, ".jpg") || hasSuffix(filename, ".jpeg"):
			c.Header("Content-Type", "image/jpeg")
		case hasSuffix(filename, ".png"):
			c.Header("Content-Type", "image/png")
		default:
			c.Header("Content-Type", "application/octet-stream")
		}

		// set the Content-Length header
		c.Header("Content-Length", strconv.FormatInt(fileSize, 10))

		// create a buffer for reading the file in chunks
		bufferSize := 1024 * 1024 * 8 // 1MB buffer size
		buffer := make([]byte, bufferSize)

		// start reading the file in chunks and writing them to the response
		for {
			bytesRead, err := file.Read(buffer)
			if err == io.EOF {
				// end of file reached
				break
			} else if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{"message": "Internal server error"})
				return
			}
			c.Writer.Write(buffer[:bytesRead])
		}
	})

	// start the server
	r.Run(":8080")
}

// helper function to check if a string has a certain suffix
func hasSuffix(str, suffix string) bool {
	return len(str) >= len(suffix) && str[len(str)-len(suffix):] == suffix
}
