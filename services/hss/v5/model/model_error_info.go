package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorInfo struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`
}

func (o ErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorInfo struct{}"
	}

	return strings.Join([]string{"ErrorInfo", string(data)}, " ")
}
