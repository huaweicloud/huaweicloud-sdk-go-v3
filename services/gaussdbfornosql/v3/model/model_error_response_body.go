package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorResponseBody struct {
	// 错误码。

	ErrorCode string `json:"error_code"`
	// 错误消息。

	ErrorMsg string `json:"error_msg"`
}

func (o ErrorResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorResponseBody struct{}"
	}

	return strings.Join([]string{"ErrorResponseBody", string(data)}, " ")
}
