package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CopyImageInRegionResponse struct {
	// 异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyImageInRegionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CopyImageInRegionResponse struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionResponse", string(data)}, " ")
}
