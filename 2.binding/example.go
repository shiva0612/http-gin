package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// User struct with validation tags
type User struct {
	Name   string `json:"name" validate:"required,min=5,max=15"`
	Age    int    `json:"age" validate:"required,min=18"`
	Mobile string `json:"mobile" validate:"required,indianmobile"`
	Email  string `json:"email" validate:"required,email"`
}

// Custom validation function for Indian mobile numbers
func validateIndianMobile(fl validator.FieldLevel) bool {
	// Extract mobile number from field value
	mobile := fl.Field().String()

	// Validate mobile number
	if len(mobile) != 10 {
		return false
	}

	if !regexp.MustCompile(`^[0-9]+$`).MatchString(mobile) {
		return false
	}

	if !regexp.MustCompile(`^(7|8|9)\d{9}$`).MatchString(mobile) {
		return false
	}

	return true
}

func test() {
	// Initialize Gin router
	r := gin.Default()

	// Initialize validator
	validate := validator.New()

	// Register custom validation function
	validate.RegisterValidation("indianmobile", validateIndianMobile)

	// POST /users endpoint
	r.POST("/users", func(c *gin.Context) {
		// Parse request body into User struct
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate User struct
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// User struct is valid, do something with it...
		// ...

		// Return success response
		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	})

	// Start server
	r.Run()
}
