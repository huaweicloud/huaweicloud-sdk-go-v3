package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckHostsResponse struct {
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckHostsResponse struct{}"
	}

	return strings.Join([]string{"CheckHostsResponse", string(data)}, " ")
}
