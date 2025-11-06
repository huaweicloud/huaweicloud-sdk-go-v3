package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Notification struct {

	// **参数解释**： 告警通知类型。 **约束限制**： 不涉及。 **取值范围**： 取值如下： notification：SMN通知； autoscaling：AS通知。 **默认取值**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 告警状态发生变化时，被通知对象的列表。 **约束限制**： 通知对象ID最多可以配置20个。 topicUrn可从SMN获取，具体操作请参考查询Topic列表。当type为notification时，notificationList列表不能为空；当type为autoscaling时，列表必须为[]。 说明：若alarm_action_enabled为true，对应的alarm_actions、insufficientdata_actions（该参数已废弃，建议无需配置）、ok_actions至少有一个不能为空。若alarm_actions、insufficientdata_actions（该参数已废弃，建议无需配置）、ok_actions同时存在时，notificationList值保持一致。
	NotificationList []string `json:"notificationList"`
}

func (o Notification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Notification struct{}"
	}

	return strings.Join([]string{"Notification", string(data)}, " ")
}
