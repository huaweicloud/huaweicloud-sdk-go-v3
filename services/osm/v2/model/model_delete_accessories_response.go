package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAccessoriesResponse struct {
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAccessoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessoriesResponse struct{}"
	}

	return strings.Join([]string{"DeleteAccessoriesResponse", string(data)}, " ")
}
