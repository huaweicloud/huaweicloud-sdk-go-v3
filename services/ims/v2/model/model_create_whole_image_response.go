package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateWholeImageResponse struct {
	// 异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWholeImageResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWholeImageResponse struct{}"
	}

	return strings.Join([]string{"CreateWholeImageResponse", string(data)}, " ")
}
