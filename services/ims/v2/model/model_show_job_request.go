package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobRequest struct {
	// 异步任务ID

	JobId string `json:"job_id"`
}

func (o ShowJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRequest", string(data)}, " ")
}
