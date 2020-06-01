package global

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var PostgressSQL *sql.DB

var InitDatabase = func() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostPSQL, portPSQL, userPSQL, passwordPSQL, dbnamePSQL)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)

		return err
	}
	PostgressSQL = db

	fmt.Println("Success connect with PSQL")

	return nil
}

var GetPSQLConn = func() *sql.DB {
	return PostgressSQL
}
