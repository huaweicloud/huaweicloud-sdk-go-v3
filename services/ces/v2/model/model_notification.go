package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Notification struct {

	// 通知类型, notification代表通过SMN通知
	Type string `json:"type"`

	// 告警状态发生变化时，被通知对象的列表。topicUrn可从SMN获取，具体操作请参考查询Topic列表。当type为notification时，notification_list列表不能为空。 说明：若alarm_action_enabled为true，对应的alarm_actions、ok_actions至少有一个不能为空。若alarm_actions、ok_actions同时存在时，notification_list值保持一致。
	NotificationList []string `json:"notification_list"`
}

func (o Notification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Notification struct{}"
	}

	return strings.Join([]string{"Notification", string(data)}, " ")
}
