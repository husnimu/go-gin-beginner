package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(service user.Service) *userHandler{
	return &userHandler{service}
}

func (h *userHandler) Signup(c *gin.Context){
	var signupRequest  user.SignupRequest

	err := c.ShouldBindJSON(&signupRequest)

	if err != nil{
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors){
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors":errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors":err.Error(),
			})
			return
		}
	}
	fmt.Println("123===========123")
	user, err := h.userService.Create(signupRequest)
	fmt.Println("1234===========1234")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors":err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":fmt.Sprintf("User %v created succesfully", user.Email),
	})
}

func (h *userHandler) Login( c *gin.Context){
	var loginRequest user.LoginRequest

	err := c.ShouldBindJSON(&loginRequest)

	if err != nil {
		switch err.(type){
		case validator.ValidationErrors:
			errorMessages := []string{}
			for _,e:= range err.(validator.ValidationErrors){
				errorMessage:= fmt.Sprintf("Error on field: %s, condition: %s", e.Field, e.ActualTag)
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, gin.H{
				"errors":err.Error(),
			})
			return
		}
	}
	tokenString, err := h.userService.Login(loginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":map[string]string{
			"token":tokenString,
		},
	})
}