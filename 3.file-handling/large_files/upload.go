package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// u can also use bufio.writer
func upload_copy_buffer() {
	// initialize a new Gin router
	r := gin.Default()

	// define a POST endpoint that accepts the form data
	r.POST("/movies", func(c *gin.Context) {
		// retrieve the form data
		movieName := c.PostForm("movie_name")
		releaseDate := c.PostForm("date_of_release")
		ratingStr := c.PostForm("rating")
		videoFile, err := c.FormFile("video")
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"message": "Bad request"})
			return
		}

		// parse the rating string to an int
		rating, err := strconv.Atoi(ratingStr)
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"message": "Bad request"})
			return
		}

		// create a new file on the server to save the video
		// using the uploaded file's original name
		fileName := fmt.Sprintf("%s-%s-%d-%s", movieName, releaseDate, rating, videoFile.Filename)
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		defer file.Close()

		// open the uploaded file for reading
		video, err := videoFile.Open()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		defer video.Close()

		// copy the contents of the uploaded file to the new file
		bufferSize := 1024 * 1024 // 1MB buffer size
		buffer := make([]byte, bufferSize)
		_, err = io.CopyBuffer(file, video, buffer)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}

		// return a success response
		c.JSON(200, gin.H{"message": "Movie saved successfully"})
	})

	// start the server
	r.Run(":8080")
}

func upload_for_loop() {
	// initialize a new Gin router
	r := gin.Default()

	// define a POST endpoint that accepts the form data
	r.POST("/movies", func(c *gin.Context) {
		// retrieve the form data
		movieName := c.PostForm("movie_name")
		releaseDate := c.PostForm("date_of_release")
		ratingStr := c.PostForm("rating")
		videoFile, err := c.FormFile("video")
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"message": "Bad request"})
			return
		}

		// parse the rating string to an int
		rating, err := strconv.Atoi(ratingStr)
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"message": "Bad request"})
			return
		}

		// create a new file on the server to save the video
		// using the uploaded file's original name
		fileName := fmt.Sprintf("%s-%s-%d-%s", movieName, releaseDate, rating, videoFile.Filename)
		file, err := os.Create(fileName)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		defer file.Close()

		// open the uploaded file for reading
		video, err := videoFile.Open()
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}
		defer video.Close()

		// copy the contents of the uploaded file to the new file
		// for{keep reading the contents using certain buffer}

		b := make([]byte, 1024*1024)
		for {
			n, err := video.Read(b)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				} else {
					c.JSON(500, gin.H{"error": "error reading the file, please try again later..."})
				}
			}
			file.Write(b[:n])

		}
		// return a success response
		c.JSON(200, gin.H{"message": "Movie saved successfully"})
	})

	// start the server
	r.Run(":8080")
}
