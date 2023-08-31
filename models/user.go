package models

type User struct {
	UserId  int      `json:"user_id"`
	Segment []string `json:"segment"`
}
