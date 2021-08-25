package models

import (
	"time"
)

type Log struct {
	Activity       string      `json:"activity"`
	URL            string      `json:"urk"`
	Request        string      `json:"request"`
	Response       interface{} `json:"response"`
	ResponseStatus int32       `json:"response_status"`
	CreatedAt      time.Time   `json:"created_at"`
}

func (b *Log) TableName() string {
	return "logs"
}
