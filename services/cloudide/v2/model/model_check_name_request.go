package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckNameRequest struct {
	// 实例名。 可以输入中文、数字、字母、下划线、点、破折号。长度介于3-100之间

	DisplayName string `json:"display_name"`
}

func (o CheckNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckNameRequest struct{}"
	}

	return strings.Join([]string{"CheckNameRequest", string(data)}, " ")
}
