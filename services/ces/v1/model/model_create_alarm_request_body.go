package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAlarmRequestBody struct {

	// **参数解释**： 告警名称。 **约束限制**： 不涉及。 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度1-128。 **默认取值**： 不涉及。
	AlarmName string `json:"alarm_name"`

	// **参数解释**： 告警描述。 **约束限制**： 不涉及。 **取值范围**： 长度[0,256]个字符。 **默认取值**： 不涉及。
	AlarmDescription *string `json:"alarm_description,omitempty"`

	Metric *CreateAlarmMetric `json:"metric"`

	Condition *Condition `json:"condition"`

	// **参数解释**： 是否启用该条告警。 **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true：开启告警。 - false：不开启告警。 **默认取值**： true
	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`

	// **参数解释**： 该条告警触发时，是否启用告警通知。 **约束限制**： 不填默认为true，对应的alarm_actions、ok_actions至少有一个不能为空。若alarm_actions、ok_actions同时存在时，alarm_actions和ok_actions中的notification_list值保持一致。 **取值范围**： 布尔值。 - true：开启告警通知。 - false：不开启告警通知。 **默认取值**： true
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`

	// **参数解释**： 告警级别。 **约束限制**： 不涉及。 **取值范围**： 只能为1、2、3、4。分别对应紧急、重要、次要、提示。 **默认取值**： 2
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// **参数解释**： 告警类型。 **约束限制**： 针对事件类型的告警时，告警类型为EVENT.SYS（系统事件）或EVENT.CUSTOM（自定义事件）。 针对资源分组的告警时，告警类型为RESOURCE_GROUP。 针对指定资源的告警时，告警类型为MULTI_INSTANCE。 **取值范围**： - EVENT.SYS：针对系统事件的告警规则。 - EVENT.CUSTOM：针对自定义事件的告警规则。 - RESOURCE_GROUP：针对资源分组的告警规则。 - MULTI_INSTANCE： 针对指定资源的告警规则。 **默认取值**： 不涉及。
	AlarmType *string `json:"alarm_type,omitempty"`

	// **参数解释**： 通知组/主题订阅的信息。 **约束限制**： 最多包含20个动作。
	AlarmActions *[]Notification `json:"alarm_actions,omitempty"`

	// **参数解释**： 通知组/主题订阅的信息。 **约束限制**： 最多包含20个动作。
	InsufficientdataActions *[]Notification `json:"insufficientdata_actions,omitempty"`

	// **参数解释**： 通知组/主题订阅的信息。 **约束限制**： 最多包含20个动作。
	OkActions *[]Notification `json:"ok_actions,omitempty"`

	// **参数解释**： 企业项目ID。如何查询企业项目ID，请参考“[获取企业项目ID](ces_03_0061.xml)”。 **约束限制**： 不涉及。 **取值范围**： 长度为0或者32个字符。 **默认取值**： 0，表示默认的企业项目default。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 告警通知开启时间。 **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。 **默认取值**： 不涉及。
	AlarmActionBeginTime *string `json:"alarm_action_begin_time,omitempty"`

	// **参数解释**： 告警通知关闭时间。 **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。 **默认取值**： 不涉及。
	AlarmActionEndTime *string `json:"alarm_action_end_time,omitempty"`
}

func (o CreateAlarmRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlarmRequestBody", string(data)}, " ")
}
