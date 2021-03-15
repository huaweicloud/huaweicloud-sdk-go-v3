package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartInstanceMasterDrActionResponse struct {
	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceMasterDrActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceMasterDrActionResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceMasterDrActionResponse", string(data)}, " ")
}
