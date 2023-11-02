package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type Person struct {
	name string
	age  string
}

func main() {

}

func using_gin(c *gin.Context) {

	router := gin.Default()

	c.Set("key", "value") //key values for this context

	c.Param("name")    // to get url params
	c.Query("name")    // to get query params
	c.QueryMap("name") // to get entire map of query params
	// c.ShouldBindJSON() for getting body

	c.PostForm("name") //value or ""
	c.FormFile("pic")  // give the first file
	c.MultipartForm()  // entire form (values map[string]string, files map[string]*fileHeaders)

	c.GetHeader("key")       //get header
	c.Header("key", "value") //set header

	c.String(400, "message", "values", "values")
	c.JSON(400, gin.H{"key": "value"})
	c.JSON(200, Person{"shiva", "24"}) //check the json tag for Person struct

	c.File("/path/filename/file.go")
	c.Data(200, "application/json", []byte("shiva"))
	// c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

	c.FullPath()
	c.Cookie("key")
	// c.SetCookie("key","value",...)
	c.Abort()                  //for aborting in middlewares
	c.AddParam("key", "value") //for internal routing
	c.ClientIP()
	c.Copy() //copy context if passing in goroutine

	w := c.Writer
	_ = w
	r := c.Request
	_ = r
	ctx := c.Request.Context()
	_ = ctx

	c.Redirect(504, "newurl")

	router.Run("8080")
	http.ListenAndServe(":8080", router)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func using_http(w http.ResponseWriter, r *http.Request) {
	// Get URL parameters
	_ = mux.Vars(r)["name"]

	// Get query parameters
	_ = r.URL.Query().Get("age")

	// BODY
	// json.NewDecoder(r.Body).Decode(bodbody)

	// Get form values
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = r.PostForm.Get("name")
	_, _, err = r.FormFile("pic")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// r.Header.Set()
	// r.Header.Get()

	json.NewEncoder(w).Encode("something in json")
}
