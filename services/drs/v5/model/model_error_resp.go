package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorResp 请求失败时返回的错误对象响应体。
type ErrorResp struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误描述。
	ErrorMsg string `json:"error_msg"`
}

func (o ErrorResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorResp struct{}"
	}

	return strings.Join([]string{"ErrorResp", string(data)}, " ")
}
