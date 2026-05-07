package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskIdRes **参数解释**： 任务ID **取值范围**: 字符长度1-64位
type TaskIdRes struct {
}

func (o TaskIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskIdRes struct{}"
	}

	return strings.Join([]string{"TaskIdRes", string(data)}, " ")
}
