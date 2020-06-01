package global

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"

	_ "github.com/lib/pq"
)

var PostgressSQL *sql.DB
var RedisConn redis.Conn

var InitDatabase = func() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostPSQL, portPSQL, userPSQL, passwordPSQL, dbnamePSQL)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	PostgressSQL = db

	fmt.Println("Success connect with PSQL")

	redisConn, err := redis.Dial(redisConnectionType, redisURL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	RedisConn = redisConn

	fmt.Println("Success connect with Redis")

	return nil
}

var GetPSQLConn = func() *sql.DB {
	return PostgressSQL
}

var GetRedisConn = func() redis.Conn {
	return RedisConn
}
