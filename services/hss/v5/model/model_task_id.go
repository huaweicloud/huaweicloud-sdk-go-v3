package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskId **参数解释**： 任务ID **取值范围**: 字符长度1-64位
type TaskId struct {
}

func (o TaskId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskId struct{}"
	}

	return strings.Join([]string{"TaskId", string(data)}, " ")
}
