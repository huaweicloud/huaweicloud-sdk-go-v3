package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ChangeTheDelayThresholdResponse struct {
	// 修改读写分离权重或延时阈值的结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeTheDelayThresholdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeTheDelayThresholdResponse struct{}"
	}

	return strings.Join([]string{"ChangeTheDelayThresholdResponse", string(data)}, " ")
}
