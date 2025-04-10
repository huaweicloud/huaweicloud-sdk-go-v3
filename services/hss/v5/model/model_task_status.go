package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskStatus 任务状态，包含如下2种   - scanning ：扫描中   - finish ：扫描完成
type TaskStatus struct {
}

func (o TaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskStatus struct{}"
	}

	return strings.Join([]string{"TaskStatus", string(data)}, " ")
}
