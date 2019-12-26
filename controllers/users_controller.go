package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/waytkheming/bookstore-users-api/domain/users"
	"github.com/waytkheming/bookstore-users-api/services"
	"github.com/waytkheming/bookstore-users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Code, err)
		return
	}

	user, err := services.GetUser(user)
	c.String(http.StatusNotImplemented, "implementme")
}

func SearchUser(c *gin.Context) {}
