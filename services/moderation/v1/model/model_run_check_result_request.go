package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunCheckResultRequest struct {
	JobId string `json:"job_id"`
}

func (o RunCheckResultRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunCheckResultRequest struct{}"
	}

	return strings.Join([]string{"RunCheckResultRequest", string(data)}, " ")
}
