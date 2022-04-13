package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 续传任务返回体
type RetryTaskResp struct {
	// 任务id

	Id string `json:"id"`
	// 状态

	Status string `json:"status"`
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o RetryTaskResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryTaskResp struct{}"
	}

	return strings.Join([]string{"RetryTaskResp", string(data)}, " ")
}
