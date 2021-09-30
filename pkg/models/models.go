package models
import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID 		int
	Title 	string
	Content string
	Created time.Time
	Expires time.Time
}

type Students struct {
	Forename string
	Surname  string
	Email string
	Password string
	Yofe int
	tutor string
	Aliveorganizationandtime string
	Aliveorganizationandtimeev string
}