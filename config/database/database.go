package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {

	var err error
	dataSourceName := "admin:mysql-todolist@tcp(database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com:3306)/db_todolist"

	// Conectar ao banco de dados MySQL
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão bem-sucedida ao banco de dados MySQL na AWS RDS!")

	return db
}
