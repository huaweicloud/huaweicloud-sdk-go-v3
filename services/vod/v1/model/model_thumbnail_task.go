package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThumbnailTask struct {

	// 任务状态，包含成功，失败
	Status *string `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	Output *ThumbnailTaskOutPut `json:"output,omitempty"`
}

func (o ThumbnailTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThumbnailTask struct{}"
	}

	return strings.Join([]string{"ThumbnailTask", string(data)}, " ")
}
