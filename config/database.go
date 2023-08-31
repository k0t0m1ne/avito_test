package config

// Данные для подключения к БД
const (
	user     = "user=sivovgg"
	password = "password=postgres_pwd"
	dbname   = "dbname=avito-test-db"
	sslmode  = "sslmode=disable"
)

const ConnStr = user + " " + password + " " + dbname + " " + sslmode
