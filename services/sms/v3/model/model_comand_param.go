package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 命令响应参数
type ComandParam struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 桶名
	Bucket *string `json:"bucket,omitempty" xml:"bucket"`
}

func (o ComandParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComandParam struct{}"
	}

	return strings.Join([]string{"ComandParam", string(data)}, " ")
}
