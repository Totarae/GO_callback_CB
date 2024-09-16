package model

import "time"

// DBObject is an object in PostgreSQL.
type DBObject struct {
	ID       uint      `gorm:"primaryKey;column:id;not null"`
	LastSeen time.Time `gorm:"column:last_seen;not null"`
}

func (DBObject) TableName() string {
	return "storage_schema.objects"
}
