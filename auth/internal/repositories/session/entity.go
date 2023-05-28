package session

import "time"

type session struct {
	Id           int64     `pg:""`
	UserId       int64     `pg:"notnull:user_id"`
	SessionToken string    `pg:"notnull:session_token"`
	ExpiresAt    time.Time `pg:""`
}
