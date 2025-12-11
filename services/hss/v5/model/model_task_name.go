package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskName **参数解释**: 任务名称 **取值范围**: 最大长度255个unicode字符。
type TaskName struct {
}

func (o TaskName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskName struct{}"
	}

	return strings.Join([]string{"TaskName", string(data)}, " ")
}
