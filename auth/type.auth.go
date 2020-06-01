package auth

type Account struct {
	UserID   int    `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

const QueryGetAccountByUsername = "SELECT * FROM account WHERE username='%s'"
const QueryInsertAccount = "INSERT INTO account (username, password) VALUES ($1, $2)"
