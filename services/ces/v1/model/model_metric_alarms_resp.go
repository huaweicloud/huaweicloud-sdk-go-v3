package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricAlarmsResp struct {

	// **参数解释**： 告警名称。 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，字符长度为[1,128]。
	AlarmName *string `json:"alarm_name,omitempty"`

	// **参数解释**： 告警描述。 **取值范围**： 长度[0,256]个字符。
	AlarmDescription *string `json:"alarm_description,omitempty"`

	// **参数解释**： 告警规则的ID或者资源分组ID。 **取值范围**： 不涉及
	RelationId *string `json:"relation_id,omitempty"`

	Metric *ListAlarmMetricResp `json:"metric,omitempty"`

	Condition *AlarmRuleConditionResp `json:"condition,omitempty"`

	// **参数解释**： 是否启用该条告警。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。
	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`

	// **参数解释**： 告警级别。 **取值范围**： 取值为1、2、3、4 - 1：紧急 - 2：重要 - 3：次要 - 4：提示
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// **参数解释**： 告警类型。 **取值范围**： 针对事件类型的告警时，告警类型为EVENT.SYS（系统事件）或EVENT.CUSTOM（自定义事件）。 针对资源分组的告警时，告警类型为RESOURCE_GROUP。 针对指定资源的告警时，告警类型为MULTI_INSTANCE。 - EVENT.SYS：针对系统事件的告警类型 - EVENT.CUSTOM：针对自定义事件的告警类型 - RESOURCE_GROUP：针对资源分组的告警类型 - MULTI_INSTANCE：指定资源的告警类型
	AlarmType *string `json:"alarm_type,omitempty"`

	// **参数解释**： 该条告警触发时，是否开启告警通知。 **取值范围**： 布尔值。 - true：开启告警通知。 - false：不开启告警通知。
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`

	// **参数解释**： 数据不足时，是否产生告警记录(不会发送通知)。 **取值范围**： 值为true或者false - true：数据不足时，产生告警记录 - false：数据不足时，不产生告警记录
	IgnoreInsufficientData *bool `json:"ignore_insufficient_data,omitempty"`

	// **参数解释**： 告警触发时，通知组/主题订阅的信息。
	AlarmActions *[]NotificationResp `json:"alarm_actions,omitempty"`

	// **参数解释**： 告警触发时，通知组/主题订阅的信息。
	OkActions *[]NotificationResp `json:"ok_actions,omitempty"`

	// **参数解释**： 告警触发时，通知组/主题订阅的信息。
	InsufficientdataActions *[]NotificationResp `json:"insufficientdata_actions,omitempty"`

	// **参数解释**： 告警规则生效的开始时间，告警规则仅在生效时间内发送通知消息。例如alarm_action_begin_time为8:00，alarm_action_end_time为20:00时，则对应的告警规则仅在08:00-20:00发送通知消息。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
	AlarmActionBeginTime *string `json:"alarm_action_begin_time,omitempty"`

	// **参数解释**： 告警规则生效的结束时间，告警规则仅在生效时间内发送通知消息。例如alarm_action_begin_time为8:00，alarm_action_end_time为20:00时，则对应的告警规则仅在08:00-20:00发送通知消息。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
	AlarmActionEndTime *string `json:"alarm_action_end_time,omitempty"`

	// **参数解释**： 时区，形如：\"GMT-08:00\"、\"GMT+08:00\"、\"GMT+0:00\" **取值范围**： 长度为[1,16]个字符。
	EffectiveTimezone *string `json:"effective_timezone,omitempty"`

	// **参数解释**： 告警规则的ID。 **取值范围**： 以al开头，后跟22个数字或字母。字符长度为24
	AlarmId *string `json:"alarm_id,omitempty"`

	// **参数解释**： 告警状态变更的时间，UNIX时间戳，单位毫秒。 **取值范围** 0 - 9999999999999
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**： 告警状态。 **取值范围**： 只能为ok、alarm、insufficient_data。 - ok：正常 - alarm：告警 - insufficient_data：数据不足
	AlarmState *string `json:"alarm_state,omitempty"`

	// **参数解释**： 企业项目ID。 **取值范围**： 只能包含小写字母、数字、“-”、“_”，长度为36个字符。也可取值0或all_granted_eps。0：代表默认企业项目ID，all_granted_eps：代表所有企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o MetricAlarmsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricAlarmsResp struct{}"
	}

	return strings.Join([]string{"MetricAlarmsResp", string(data)}, " ")
}
