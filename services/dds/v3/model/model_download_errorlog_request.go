package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DownloadErrorlogRequest struct {
	InstanceId string                       `json:"instance_id"`
	Body       *DownloadErrorlogRequestBody `json:"body,omitempty"`
}

func (o DownloadErrorlogRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DownloadErrorlogRequest struct{}"
	}

	return strings.Join([]string{"DownloadErrorlogRequest", string(data)}, " ")
}
