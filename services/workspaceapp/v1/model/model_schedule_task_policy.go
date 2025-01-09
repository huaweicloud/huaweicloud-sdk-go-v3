package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScheduleTaskPolicy 定时任务策略。
type ScheduleTaskPolicy struct {

	// 当存在会话的时候，是否强制执行，强制执行开关。取值为： * false：表示关闭。 * true：表示开启。
	EnforcementEnable *bool `json:"enforcement_enable,omitempty"`
}

func (o ScheduleTaskPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleTaskPolicy struct{}"
	}

	return strings.Join([]string{"ScheduleTaskPolicy", string(data)}, " ")
}
