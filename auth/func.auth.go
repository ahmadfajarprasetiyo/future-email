package auth

import (
	"fmt"

	"../global"
	"github.com/gin-gonic/gin"
)

var Login = func(c *gin.Context) {
	var account Account

	username := c.PostForm("username")
	// password := c.PostForm("password")

	psql := global.GetPSQLConn()

	query := fmt.Sprintf(QueryGetAccountByUsername, username)

	err := psql.QueryRow(query).Scan(&account.UserID, &account.Username, &account.Password)

	if err == nil {
		c.JSON(200, gin.H{
			"status":   "success",
			"password": account.Password,
		})
	} else {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":   "failed",
			"password": "-",
		})
	}
}

var Register = func(c *gin.Context) {

}
