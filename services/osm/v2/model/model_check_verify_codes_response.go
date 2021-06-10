package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CheckVerifyCodesResponse struct {
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckVerifyCodesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckVerifyCodesResponse struct{}"
	}

	return strings.Join([]string{"CheckVerifyCodesResponse", string(data)}, " ")
}
