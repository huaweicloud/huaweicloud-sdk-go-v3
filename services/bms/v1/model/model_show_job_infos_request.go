package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobInfosRequest struct {
	// Job ID

	JobId string `json:"job_id"`
}

func (o ShowJobInfosRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobInfosRequest struct{}"
	}

	return strings.Join([]string{"ShowJobInfosRequest", string(data)}, " ")
}
