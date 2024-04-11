package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpgradeErrorResponseBody struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o UpgradeErrorResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeErrorResponseBody struct{}"
	}

	return strings.Join([]string{"UpgradeErrorResponseBody", string(data)}, " ")
}
