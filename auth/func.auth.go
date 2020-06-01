package auth

import (
	"database/sql"
	"fmt"

	"../global"
	"../utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var getAccountByUsername = func(username string) (Account, error) {
	var account Account

	psql := global.GetPSQLConn()
	query := fmt.Sprintf(QueryGetAccountByUsername, username)
	err := psql.QueryRow(query).Scan(&account.UserID, &account.Username, &account.Password)

	return account, err
}

var generatePasswordHash = func(password string) (string, error) {
	passwordByte := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.MinCost)

	return string(passwordHash), err
}

var generateToken = func(userID int) (string, error) {
	randomString := utils.RandStringBytes(LengthToken)

	redisConn := global.GetRedisConn()
	keyRedis := fmt.Sprintf(KeyRedisToken, userID)

	_, err := redisConn.Do("SETEX", keyRedis, global.ExpiredTimeToken, randomString)
	if err != nil {
		fmt.Println(err)
		return "-", err
	}

	return randomString, nil
}

var Login = func(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	account, err := getAccountByUsername(username)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": "failed",
			"token":  "-",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": "failed",
			"token":  "-",
		})
		return
	}

	token, err := generateToken(account.UserID)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": "failed",
			"token":  "-",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"token":  token,
	})
	return
}

var Register = func(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	_, err := getAccountByUsername(username)

	if err != sql.ErrNoRows {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	password, err = generatePasswordHash(password)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	psql := global.GetPSQLConn()
	_, err = psql.Exec(QueryInsertAccount, username, password)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
	})

	return
}
