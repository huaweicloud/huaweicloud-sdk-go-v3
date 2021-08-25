package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestoreInstanceResponse struct {
	// 恢复到当前实例的异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreInstanceResponse", string(data)}, " ")
}
