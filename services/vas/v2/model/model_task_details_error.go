package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业运行失败时收集到的错误信息
type TaskDetailsError struct {
	// 算法服务定义的错误码

	Code string `json:"code"`
	// 算法服务反馈的错误信息

	Message string `json:"message"`
}

func (o TaskDetailsError) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDetailsError struct{}"
	}

	return strings.Join([]string{"TaskDetailsError", string(data)}, " ")
}
