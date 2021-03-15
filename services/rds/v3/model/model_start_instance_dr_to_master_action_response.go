package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartInstanceDrToMasterActionResponse struct {
	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceDrToMasterActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceDrToMasterActionResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceDrToMasterActionResponse", string(data)}, " ")
}
