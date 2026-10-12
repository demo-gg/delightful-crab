package main

import "time"

type UploadRequest struct {
	Key         string
	Data        []byte
	ContentType string
	Metadata    map[string]string
}

type UploadResponse struct {
	ObjectKey string
	ETag      string
	VersionID string
	Location  string
}

type UploadMetadata struct {
	Key       string
	Size      int64
	Timestamp time.Time
	Success   bool
	Error     string
}

func NewUploadRequest(key string, data []byte) *UploadRequest {
	return &UploadRequest{
		Key:  key,
		Data: data,
	}
}
