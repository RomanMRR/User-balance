package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)


func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id not found")
	}
		

	return idInt, nil
}