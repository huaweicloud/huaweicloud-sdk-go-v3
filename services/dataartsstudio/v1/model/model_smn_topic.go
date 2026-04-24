package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SmnTopic struct {

	// 名称。
	Name *string `json:"name,omitempty"`

	// smn urn。
	Urn *string `json:"urn,omitempty"`

	// 告警方式：主题、责任人，值班表和钉钉机器人。
	NotifyMethod *SmnTopicNotifyMethod `json:"notify_method,omitempty"`

	// 告警协议：短信、邮件、电话，示例：[\"email\"]  \"email\"：邮件传输协议,\"sms\"：短信传输协议,\"callnotify\":语音, \"dingding\":个人钉钉。
	Protocol *string `json:"protocol,omitempty"`

	// 抄送人，示例：[\"lin\",\"hua\"]。
	OtherPersons *string `json:"other_persons,omitempty"`

	// 最大告警次数。
	MaxSendTimes *int32 `json:"max_send_times,omitempty"`

	// 告警间隔。
	SendInterval *int32 `json:"send_interval,omitempty"`

	// 值班表id。
	DutyScheduleId *int64 `json:"duty_schedule_id,omitempty"`

	// 值班表名称。
	DutyScheduleName *string `json:"duty_schedule_name,omitempty"`

	// 机器人id。
	SmnConfigId *string `json:"smn_config_id,omitempty"`

	// 机器人名称。
	SmnConfigName *string `json:"smn_config_name,omitempty"`
}

func (o SmnTopic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmnTopic struct{}"
	}

	return strings.Join([]string{"SmnTopic", string(data)}, " ")
}

type SmnTopicNotifyMethod struct {
	value string
}

type SmnTopicNotifyMethodEnum struct {
	SUBJECT       SmnTopicNotifyMethod
	OWNER         SmnTopicNotifyMethod
	DUTY_SCHEDULE SmnTopicNotifyMethod
	SUBJECT_OWNER SmnTopicNotifyMethod
	SUBJECT_DUTY  SmnTopicNotifyMethod
	DINGDING      SmnTopicNotifyMethod
}

func GetSmnTopicNotifyMethodEnum() SmnTopicNotifyMethodEnum {
	return SmnTopicNotifyMethodEnum{
		SUBJECT: SmnTopicNotifyMethod{
			value: "SUBJECT",
		},
		OWNER: SmnTopicNotifyMethod{
			value: "OWNER",
		},
		DUTY_SCHEDULE: SmnTopicNotifyMethod{
			value: "DUTY_SCHEDULE",
		},
		SUBJECT_OWNER: SmnTopicNotifyMethod{
			value: "SUBJECT_OWNER",
		},
		SUBJECT_DUTY: SmnTopicNotifyMethod{
			value: "SUBJECT_DUTY",
		},
		DINGDING: SmnTopicNotifyMethod{
			value: "DINGDING",
		},
	}
}

func (c SmnTopicNotifyMethod) Value() string {
	return c.value
}

func (c SmnTopicNotifyMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmnTopicNotifyMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
