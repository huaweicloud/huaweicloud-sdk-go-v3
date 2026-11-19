package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRemindersConfigRequestInfo 设置告警提醒开关请求体
type SetRemindersConfigRequestInfo struct {

	// **参数解释**: 是否开启告警提醒 **约束限制**: 必填 **取值范围**: - true：开启 - false：关闭  **默认取值**: 不涉及
	Reminders bool `json:"reminders"`
}

func (o SetRemindersConfigRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRemindersConfigRequestInfo struct{}"
	}

	return strings.Join([]string{"SetRemindersConfigRequestInfo", string(data)}, " ")
}
