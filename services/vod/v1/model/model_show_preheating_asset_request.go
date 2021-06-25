package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPreheatingAssetRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
}

func (o ShowPreheatingAssetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPreheatingAssetRequest struct{}"
	}

	return strings.Join([]string{"ShowPreheatingAssetRequest", string(data)}, " ")
}
