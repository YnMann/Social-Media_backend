package models

import "time"

type Messages struct {
	Id                int
	Text              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Author_user_id    int
	Recipient_user_id int
	Is_read           bool
}
