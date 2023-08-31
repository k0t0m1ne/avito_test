package initializer

import (
	"avitoTest/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func DbConnectionInit() *sql.DB {
	// Инициализируем подключение к БД, используя данные из database.go
	db, err := sql.Open("postgres", config.ConnStr)
	if err != nil {
		log.Println(err)
		return nil
	} else {
		log.Printf("Connection to database opened successfully")
	}
	return db
}
