package models

import "time"

type Messages struct {
	ID        string
	Sender    string
	Recipient string
	Content   string
	ServerIP  string
	SenderIP  string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsRead    bool
}
