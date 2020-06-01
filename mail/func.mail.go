package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"../utils"
	"github.com/gin-gonic/gin"
)

var SendMailToNode = func(c *gin.Context) {
	email := c.PostForm("email")
	timeSendString := c.PostForm("time_send")
	content := c.PostForm("content")

	if !utils.IsValidEmail(email) {
		fmt.Println("Email is invalid")
		utils.ErrorResponse(c)
		return
	}

	timeSend, err := strconv.ParseInt(timeSendString, 10, 64)
	if err != nil {
		fmt.Println(err)
		utils.ErrorResponse(c)
		return
	}

	if !utils.IsValidTime(timeSend) {
		fmt.Println("Time send is invalid")
		utils.ErrorResponse(c)
		return
	}

	sendEmailRequest := SendEmailRequest{
		Email:    email,
		Content:  content,
		TimeSend: timeSend,
	}

	b, _ := json.Marshal(sendEmailRequest)
	uri, err := url.Parse(URLNodeSendEmail)

	req, err := http.NewRequest(http.MethodPost, uri.String(), bytes.NewBuffer(b))

	if err != nil {
		fmt.Println(err)
		utils.ErrorResponse(c)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	timeout := time.Duration(45 * time.Second)

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		utils.ErrorResponse(c)
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		utils.ErrorResponse(c)
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
	})

	return
}
