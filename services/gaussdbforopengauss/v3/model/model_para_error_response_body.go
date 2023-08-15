package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParaErrorResponseBody struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o ParaErrorResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParaErrorResponseBody struct{}"
	}

	return strings.Join([]string{"ParaErrorResponseBody", string(data)}, " ")
}
