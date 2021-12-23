package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAuthorizationsResponse struct {
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuthorizationsResponse", string(data)}, " ")
}
