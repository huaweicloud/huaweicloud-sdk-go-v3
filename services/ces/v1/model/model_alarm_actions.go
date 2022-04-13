package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmActions struct {
	// 告警通知类型，取值如下： notification：通知； autoscaling：弹性伸缩。

	Type string `json:"type"`
	// 告警状态发生变化时，被通知对象的列表。通知对象ID最多可以配置5个。topicUrn可从SMN获取，具体操作请参考查询Topic列表。当type为notification时，notificationList列表不能为空；当type为autoscaling时，列表必须为[]。 说明：若alarm_action_enabled为true，对应的alarm_actions、insufficientdata_actions（该参数已废弃，建议无需配置）、ok_actions至少有一个不能为空。若alarm_actions、insufficientdata_actions（该参数已废弃，建议无需配置）、ok_actions同时存在时，notificationList值保持一致。

	NotificationList []string `json:"notificationList"`
}

func (o AlarmActions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmActions struct{}"
	}

	return strings.Join([]string{"AlarmActions", string(data)}, " ")
}
