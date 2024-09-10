package model

import "time"

// DBObject is an object in PostgreSQL.
type DBObject struct {
	ID        uint      `pg:"id,notnull,pk"`
	LastSeen  time.Time `pg:"last_seen,notnull"`
	tableName struct{}  `pg:"storage_schema.objects"`
}
