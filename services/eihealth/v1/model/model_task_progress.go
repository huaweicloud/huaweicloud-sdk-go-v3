package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskProgress 任务进度
type TaskProgress struct {

	// 整体进度
	Overall *float32 `json:"overall,omitempty"`
}

func (o TaskProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskProgress struct{}"
	}

	return strings.Join([]string{"TaskProgress", string(data)}, " ")
}
