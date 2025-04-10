package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskId 任务ID
type TaskId struct {
}

func (o TaskId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskId struct{}"
	}

	return strings.Join([]string{"TaskId", string(data)}, " ")
}
