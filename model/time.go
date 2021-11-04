package model

import (
	"time"
)

type AccessInfo struct {
	Timestamp	time.Time	`json:"timestamp"`
	Latency		int64		`json:"latenInfo"`
	Path		string		`json:"path"`
	OS			string		`json:"os"`
}

func (ap *AccessInfo) StoreTimestamp(timestamp time.Time) *AccessInfo {
	ap.Timestamp = timestamp
	return ap
}

func (ap *AccessInfo) StoreLatency(latency int64) *AccessInfo {
	ap.Latency = latency
	return ap
}

func (ap *AccessInfo) StorePath(path string) *AccessInfo {
	ap.Path = path
	return ap
}

func (ap *AccessInfo) StoreOS(os string) *AccessInfo {
	ap.OS = os
	return ap
}
