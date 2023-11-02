package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// simple()
	complex()
}
func simple() {
	router := gin.Default()
	router.Static("/assets", "../assets") //serving the folder, localhost:8080/assets/file1.txt

	router.StaticFile("/file1", "../assets/file.go")

	router.GET("/file", func(c *gin.Context) {
		c.File("../assets/file.go")
	})

	router.Run(":8080")

}

// streaming a video or large file
func complex() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, _ := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")

		reader := response.Body
		defer reader.Close()
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.Run(":8080")
}

//to display file in browser - Content-Disposition: inline
//save as dialogue with defaultfilename - Content-Disposition: attachment
//save as dialogue with custom filename - Content-Disposition: attachment; filename="filename.jpg"
