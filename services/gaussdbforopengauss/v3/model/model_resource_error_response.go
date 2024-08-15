package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceErrorResponse struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o ResourceErrorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceErrorResponse struct{}"
	}

	return strings.Join([]string{"ResourceErrorResponse", string(data)}, " ")
}
