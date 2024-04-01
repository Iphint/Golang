package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang-test")
	if err != nil {
		fmt.Println("Error connecting to database:", err.Error())
		return
	}
	DB = db
	fmt.Println("Connection successfully")
}
