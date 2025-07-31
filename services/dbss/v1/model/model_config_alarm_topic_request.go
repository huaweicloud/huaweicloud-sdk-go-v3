package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigAlarmTopicRequest struct {

	// CPU告警阈值(%)
	AlarmCpu int32 `json:"alarm_cpu"`

	// 磁盘告警阈值(%)
	AlarmDisk int32 `json:"alarm_disk"`

	// 内存告警阈值(%)
	AlarmMemory int32 `json:"alarm_memory"`

	// 每天发送告警总条数
	AlarmNum int32 `json:"alarm_num"`

	// 告警等级,默认为空。 - high：高  - medium：中  - low：低
	AlarmRisk *[]string `json:"alarm_risk,omitempty"`

	// 通知开关 - ON：开启  - OFF：关闭
	AlarmSwitch string `json:"alarm_switch"`

	// 通知消息主题URN,调用SMN服务接口获取。当alarm_switch为ON时必填
	AlarmTopicUrn *string `json:"alarm_topic_urn,omitempty"`
}

func (o ConfigAlarmTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigAlarmTopicRequest struct{}"
	}

	return strings.Join([]string{"ConfigAlarmTopicRequest", string(data)}, " ")
}
