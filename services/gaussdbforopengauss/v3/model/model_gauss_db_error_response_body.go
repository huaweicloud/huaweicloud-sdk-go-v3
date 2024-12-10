package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GaussDbErrorResponseBody struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o GaussDbErrorResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GaussDbErrorResponseBody struct{}"
	}

	return strings.Join([]string{"GaussDbErrorResponseBody", string(data)}, " ")
}
