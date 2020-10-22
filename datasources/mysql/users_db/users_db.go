package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	//Client is
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	//datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s","admin", "pass123", "127.0.0.1:3306", "users_db")
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema) //setting env variable becz not secure to add this code in github
	//fmt.Println("---->>", datasourceName)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		fmt.Println("err---->>", err)
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database sucessfully connected")
}
