package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExportImageResponse struct {
	// 异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportImageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportImageResponse struct{}"
	}

	return strings.Join([]string{"ExportImageResponse", string(data)}, " ")
}
