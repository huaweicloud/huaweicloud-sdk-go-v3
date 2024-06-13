package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiError 失败时的error信息
type ApiError struct {

	// 业务失败的错误码
	Code *string `json:"code,omitempty"`

	// 业务失败的提示内容
	Reason *string `json:"reason,omitempty"`
}

func (o ApiError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiError struct{}"
	}

	return strings.Join([]string{"ApiError", string(data)}, " ")
}
