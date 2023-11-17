package database

import (
	"database/sql"
	"log"

	"github.com/eron97/crud_nov.git/config/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {

	var err error
	dataSourceName := "admin:mysql-todolist@tcp(database-1.cpj0eavfzshu.us-east-1.rds.amazonaws.com:3306)/db_todolist"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[InitDB: Conexão bem-sucedida ao AWS RDS]")

	return db
}

// GetDB retorna o ponteiro do banco de dados do contexto Gin
func GetDB(c *gin.Context) *sql.DB {
	db, exists := c.Get("db")
	if !exists {
		log.Fatal("[GetDB: Ponteiro não está no contexto Gin]")
	}
	return db.(*sql.DB)
}

func QueryAllTasks(db *sql.DB) ([]models.Task, error) {

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Task_Name, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetTaskByName retorna uma tarefa específica pelo nome
func QueryByName(db *sql.DB, taskName string) (*models.Task, error) {
	row := db.QueryRow("SELECT * FROM tasks WHERE Task_Name = ?", taskName)

	var task models.Task
	err := row.Scan(&task.ID, &task.Task_Name, &task.Priority)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &task, nil
}

// AddTask adiciona uma nova tarefa ao banco de dados
func QueryAddTask(db *sql.DB, task *models.TaskPost) error {
	_, err := db.Exec("INSERT INTO tasks (Task_Name, Priority) VALUES (?, ?)", task.Task_Name, task.Priority)
	if err != nil {
		log.Printf("Erro na tarefa do package database: %s ", err.Error())
		return err
	}
	return nil
}
