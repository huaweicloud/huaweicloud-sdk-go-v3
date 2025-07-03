package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduledEventUpdateBody 修改计划事件计划执行开始时间
type ScheduledEventUpdateBody struct {

	// 计划执行开始时间，新的开始时间必须早于not_before_deadline
	NotBefore string `json:"not_before"`
}

func (o ScheduledEventUpdateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledEventUpdateBody struct{}"
	}

	return strings.Join([]string{"ScheduledEventUpdateBody", string(data)}, " ")
}
