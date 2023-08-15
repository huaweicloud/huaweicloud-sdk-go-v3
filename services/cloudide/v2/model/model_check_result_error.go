package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckResultError struct {

	// 解析结果码
	Code *string `json:"code,omitempty"`

	// 解析结果信息
	Message *string `json:"message,omitempty"`
}

func (o CheckResultError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckResultError struct{}"
	}

	return strings.Join([]string{"CheckResultError", string(data)}, " ")
}
