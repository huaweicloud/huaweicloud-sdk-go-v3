package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Error 提交任务异常时返回的异常信息
type Error struct {

	// 任务异常错误信息描述
	Message *string `json:"message,omitempty"`

	// 任务异常错误信息编码
	Code *string `json:"code,omitempty"`
}

func (o Error) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Error struct{}"
	}

	return strings.Join([]string{"Error", string(data)}, " ")
}
