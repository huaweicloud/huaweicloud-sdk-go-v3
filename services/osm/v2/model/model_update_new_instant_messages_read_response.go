package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNewInstantMessagesReadResponse struct{}"
	}

	return strings.Join([]string{"UpdateNewInstantMessagesReadResponse", string(data)}, " ")
}
