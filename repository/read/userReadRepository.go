package read

import (
	"avitoTest/models"
	"database/sql"
	"log"
)

func GetUserById(db *sql.DB, id int) models.User {

	var u = models.User{}
	u.UserId = id
	var segment string
	var segments []string

	rows, _ := db.Query("select s.seg_name from user_segment join segments s on user_segment.segment = s.id where user_id = $1 ", id)

	for rows.Next() {
		err := rows.Scan(&segment)
		if err != nil {
			log.Fatal("data is not correct")
		}
		segments = append(segments, segment)
	}
	u.Segment = segments
	return u
}
