package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorInfo struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误详情
	ErrorDetail *string `json:"error_detail,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorInfo struct{}"
	}

	return strings.Join([]string{"ErrorInfo", string(data)}, " ")
}
