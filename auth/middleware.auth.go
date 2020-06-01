package auth

import (
	"fmt"
	"strconv"

	"../utils"
	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		userIDString := c.GetHeader("X-UserID")
		userID, err := strconv.Atoi(userIDString)
		if err != nil {
			fmt.Println(err)
			utils.ErrorResponse(c)
			c.Abort()
			return
		}

		token := c.GetHeader("Authorization")

		valid := checkIsValidToken(userID, token)
		if !valid {
			utils.ErrorResponse(c)
			c.Abort()
			return
		}

		c.Next()
		return
	}
}
