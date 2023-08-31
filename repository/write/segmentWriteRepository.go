package write

import (
	"database/sql"
	"log"
)

func AddSegment(db *sql.DB, segName string) {
	_, err := db.Exec("INSERT INTO segments(seg_name) VALUES ($1)", segName)
	if err != nil {
		log.Println("error while inserting data")
	}
}

func DeleteSegment(db *sql.DB, segName string) {
	_, err := db.Exec("DELETE FROM segments WHERE seg_name = $1", segName)
	if err != nil {
		log.Println("error while deleting data")
	}
}
