package mail

type SendEmailRequest struct {
	Email    string `json:"email"`
	Content  string `json:"content"`
	TimeSend int64  `json:"time_send"`
}

const URLNodeSendEmail = "http://localhost:8081/send"
