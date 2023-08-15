package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComandParam 命令响应参数
type ComandParam struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 桶名
	Bucket *string `json:"bucket,omitempty"`
}

func (o ComandParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComandParam struct{}"
	}

	return strings.Join([]string{"ComandParam", string(data)}, " ")
}
