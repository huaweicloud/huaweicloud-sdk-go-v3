package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ConfirmAuthorizationsResponse struct {
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ConfirmAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"ConfirmAuthorizationsResponse", string(data)}, " ")
}
