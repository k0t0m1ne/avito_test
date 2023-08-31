package write

import "database/sql"

func AddUserToSegment(db *sql.DB, segmentsToAdd []string, SegmentsToDelete []string, id int) {
	if len(segmentsToAdd) != 0 {
		for i := range segmentsToAdd {
			db.Exec("INSERT INTO user_segment(user_id, segment) VALUES ($1, (select id from segments where seg_name = $2))",
				id, segmentsToAdd[i])
		}
	}

	if len(SegmentsToDelete) != 0 {
		for i := range SegmentsToDelete {
			db.Exec("DELETE FROM user_segment WHERE user_id = $1 AND segment = (select id from segments where seg_name = $2)",
				id, SegmentsToDelete[i])
		}
	}
}
