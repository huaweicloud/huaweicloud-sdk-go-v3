package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateMessagesResponse struct {
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessagesResponse struct{}"
	}

	return strings.Join([]string{"CreateMessagesResponse", string(data)}, " ")
}
