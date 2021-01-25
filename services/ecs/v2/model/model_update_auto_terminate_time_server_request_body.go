package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateAutoTerminateTimeServerRequestBody struct {
	// 销毁时间
	AutoTerminateTime string `json:"auto_terminate_time"`
}

func (o UpdateAutoTerminateTimeServerRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAutoTerminateTimeServerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAutoTerminateTimeServerRequestBody", string(data)}, " ")
}
