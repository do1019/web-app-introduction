package middleware

import (
	"time"
)

type AccessInfo struct {
	Timestamp	time.Time	`json:"timestamp"`
	Latency		int64		`json:"latenInfo"`
	Path		string		`json:"path"`
	OS			string		`json:"os"`
}

func (a *AccessInfo) StoreTimestamp(timestamp time.Time) *AccessInfo {
	a.Timestamp = timestamp
	return a
}

func (a *AccessInfo) StoreLatency(latency int64) *AccessInfo {
	a.Latency = latency
	return a
}

func (a *AccessInfo) StorePath(path string) *AccessInfo {
	a.Path = path
	return a
}

func (a *AccessInfo) StoreOS(os string) *AccessInfo {
	a.OS = os
	return a
}
