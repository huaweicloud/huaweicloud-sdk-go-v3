package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateServerAutoTerminateTimeRequestBody struct {
	// 销毁时间

	AutoTerminateTime string `json:"auto_terminate_time"`
}

func (o UpdateServerAutoTerminateTimeRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateServerAutoTerminateTimeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServerAutoTerminateTimeRequestBody", string(data)}, " ")
}
