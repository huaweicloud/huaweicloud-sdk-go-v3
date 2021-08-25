package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobDetailRequest struct {
	// 任务ID。

	Id string `json:"id"`
}

func (o ShowJobDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJobDetailRequest", string(data)}, " ")
}
