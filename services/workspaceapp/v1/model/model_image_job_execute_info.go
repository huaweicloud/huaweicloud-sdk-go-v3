package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageJobExecuteInfo 任务执行失败信息
type ImageJobExecuteInfo struct {

	// 任务执行失败时的错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 任务失败原因。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ImageJobExecuteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageJobExecuteInfo struct{}"
	}

	return strings.Join([]string{"ImageJobExecuteInfo", string(data)}, " ")
}
