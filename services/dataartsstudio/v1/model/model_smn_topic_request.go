package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmnTopicRequest struct {

	// 主题名称。
	Name string `json:"name"`

	// smn urn，可以在SMN服务查看对应主题的urn。
	Urn string `json:"urn"`

	// 告警方式：主题、责任人、值班表和钉钉机器人，取值如下： - SUBJECT: 主题 - OWNER： 责任人 - DUTY_SCHEDULE: 值班表 - DINGDING: 钉钉机器人
	NotifyMethod string `json:"notify_method"`

	// 告警协议：短信、邮件、电话，示例：[\"email\"] \"email\"：邮件传输协议,\"sms\"：短信传输协议,\"callnotify\":语音, \"dingding\":个人钉钉。
	Protocol *string `json:"protocol,omitempty"`

	// 抄送人，示例：[\"lin\",\"hua\"]。
	OtherPersons *string `json:"other_persons,omitempty"`

	// 最大告警次数，取值为 [1, 50]。
	MaxSendTimes int32 `json:"max_send_times"`

	// 告警间隔，取值为[5, 60]，单位：分钟。
	SendInterval int32 `json:"send_interval"`

	// 值班表名称。
	DutyScheduleName *string `json:"duty_schedule_name,omitempty"`

	// 机器人名称。
	SmnConfigName *string `json:"smn_config_name,omitempty"`
}

func (o SmnTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnTopicRequest struct{}"
	}

	return strings.Join([]string{"SmnTopicRequest", string(data)}, " ")
}
