package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateNewInstantMessagesReadResponse struct {
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNewInstantMessagesReadResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateNewInstantMessagesReadResponse struct{}"
	}

	return strings.Join([]string{"UpdateNewInstantMessagesReadResponse", string(data)}, " ")
}
