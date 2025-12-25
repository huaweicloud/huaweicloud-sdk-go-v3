package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeRunInfoInnerErrorErrorBody struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o NodeRunInfoInnerErrorErrorBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeRunInfoInnerErrorErrorBody struct{}"
	}

	return strings.Join([]string{"NodeRunInfoInnerErrorErrorBody", string(data)}, " ")
}
