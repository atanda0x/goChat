package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    string
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	rows, err := db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}
