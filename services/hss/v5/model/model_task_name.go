package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskName 任务名称
type TaskName struct {
}

func (o TaskName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskName struct{}"
	}

	return strings.Join([]string{"TaskName", string(data)}, " ")
}
