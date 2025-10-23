package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateActionRules struct {

	// 告警规则id。
	AlarmRuleId int64 `json:"alarm_rule_id"`

	// 告警规则名称。
	AlarmRuleName string `json:"alarm_rule_name"`

	// 告警规则类型。
	AlarmRuleType string `json:"alarm_rule_type"`

	// 待绑定的告警行动规则名称。
	BindNotificationRuleId *string `json:"bind_notification_rule_id,omitempty"`

	// 通知频率 - 当通知类型为“alarm_policy”时，填“-1” - 当通知类型为“direct”时，    - “0”：只告警一次    - “300”：每5分钟    - “600”：每10分钟    - “900”：每15分钟    - “1800”：每30分钟    - “3600”：每1小时    - “10800”：每3小时    - “21600”：每6小时    - “43200”：每12小时    - “86400”：每天
	Frequency string `json:"frequency"`

	// 是否启用告警通知规则。 - 当通知类型为“direct”时，填true - 当通知类型为“alarm_policy”时，填false 如果告警触发“notify_triggered”或告警恢复“notify_resolved”都设置为false（即都不进行告警通知），则notification_enable需设置为false。
	NotificationEnable *bool `json:"notification_enable,omitempty"`

	// 通知类型。 - “direct”：直接告警 - “alarm_policy”：告警降噪
	NotificationType *string `json:"notification_type,omitempty"`

	// 告警解决是否通知。 - true：通知 - false：不通知
	NotifyResolved bool `json:"notify_resolved"`

	// 告警触发是否通知。 - true：通知 - false：不通知
	NotifyTriggered bool `json:"notify_triggered"`

	// 启用告警分组规则。 - 当通知类型为“alarm_policy”时：true - 当通知类型为“direct”时：false 如果告警触发“notify_triggered”或告警恢复“notify_resolved”都设置为false（即都不进行告警通知），则route_group_enable需设置为false。
	RouteGroupEnable *bool `json:"route_group_enable,omitempty"`

	// 告警分组规则名称。 - 当route_group_enable 为true时，填告警分组规则名称 - 当route_group_enable 为false时，填“”
	RouteGroupRule *string `json:"route_group_rule,omitempty"`
}

func (o BatchUpdateActionRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateActionRules struct{}"
	}

	return strings.Join([]string{"BatchUpdateActionRules", string(data)}, " ")
}
