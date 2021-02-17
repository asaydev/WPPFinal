package model

import "time"

type Session struct {
	Un           string
	LastActivity time.Time
}
