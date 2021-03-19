package model

import (
	"encoding/json"

	"strings"
)

// 错误码消息
type ErrorInfo struct {
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ErrorInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ErrorInfo struct{}"
	}

	return strings.Join([]string{"ErrorInfo", string(data)}, " ")
}
