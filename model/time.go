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

func (ap *AccessInfo) withTimestamp(timestamp time.Time) *AccessInfo {
	ap.Timestamp = timestamp
	return ap
}

func (ap *AccessInfo) withLatency(latency int64) *AccessInfo {
	ap.Latency = latency
	return ap
}

func (ap *AccessInfo) withPath(path string) *AccessInfo {
	ap.Path = path
	return ap
}

func (ap *AccessInfo) withOS(os string) *AccessInfo {
	ap.OS = os
	return ap
}
