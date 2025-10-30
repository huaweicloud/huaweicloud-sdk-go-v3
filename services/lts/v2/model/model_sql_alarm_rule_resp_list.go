package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SqlAlarmRuleRespList struct {

	// SQL告警名称
	SqlAlarmRuleName string `json:"sql_alarm_rule_name"`

	// 是否管道符sql查询
	IsCssSql *bool `json:"is_css_sql,omitempty"`

	// SQL告警规则id
	SqlAlarmRuleId string `json:"sql_alarm_rule_id"`

	// SQL告警信息描述
	SqlAlarmRuleDescription string `json:"sql_alarm_rule_description"`

	// SQL详细信息
	SqlRequests []SqlRequest `json:"sql_requests"`

	Frequency *FrequencyRespBody `json:"frequency"`

	// 条件表达式
	ConditionExpression string `json:"condition_expression"`

	// 主题信息
	Topics []Topics `json:"topics"`

	// 告警级别
	SqlAlarmLevel SqlAlarmRuleRespListSqlAlarmLevel `json:"sql_alarm_level"`

	// domainId
	DomainId string `json:"domain_id"`

	// 创建时间（毫秒时间戳）
	CreateTime int64 `json:"create_time"`

	// 更新时间（毫秒时间戳）
	UpdateTime int64 `json:"update_time"`

	// 消息模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 告警状态
	Status *SqlAlarmRuleRespListStatus `json:"status,omitempty"`

	// 触发条件：触发周期;默认为1
	TriggerConditionCount *int32 `json:"trigger_condition_count,omitempty"`

	// 触发条件：触发周期;默认为1
	TriggerConditionFrequency *int32 `json:"trigger_condition_frequency,omitempty"`

	// 是否打开恢复通知;默认false
	WhetherRecoveryPolicy *bool `json:"whether_recovery_policy,omitempty"`

	// 恢复策略周期;默认为3
	RecoveryPolicy *int32 `json:"recovery_policy,omitempty"`

	// 通知频率,单位(分钟)
	NotificationFrequency SqlAlarmRuleRespListNotificationFrequency `json:"notification_frequency"`

	// 告警行动规则名称 >alarm_action_rule_name和notification_save_rule可以选填一个，如果都填，优先选择alarm_action_rule_name
	AlarmActionRuleName *string `json:"alarm_action_rule_name,omitempty"`

	// **参数解释：** 告警标签信息。
	Tags *[]TagsResBody `json:"tags,omitempty"`
}

func (o SqlAlarmRuleRespList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlAlarmRuleRespList struct{}"
	}

	return strings.Join([]string{"SqlAlarmRuleRespList", string(data)}, " ")
}

type SqlAlarmRuleRespListSqlAlarmLevel struct {
	value string
}

type SqlAlarmRuleRespListSqlAlarmLevelEnum struct {
	INFO     SqlAlarmRuleRespListSqlAlarmLevel
	MINOR    SqlAlarmRuleRespListSqlAlarmLevel
	MAJOR    SqlAlarmRuleRespListSqlAlarmLevel
	CRITICAL SqlAlarmRuleRespListSqlAlarmLevel
}

func GetSqlAlarmRuleRespListSqlAlarmLevelEnum() SqlAlarmRuleRespListSqlAlarmLevelEnum {
	return SqlAlarmRuleRespListSqlAlarmLevelEnum{
		INFO: SqlAlarmRuleRespListSqlAlarmLevel{
			value: "Info",
		},
		MINOR: SqlAlarmRuleRespListSqlAlarmLevel{
			value: "Minor",
		},
		MAJOR: SqlAlarmRuleRespListSqlAlarmLevel{
			value: "Major",
		},
		CRITICAL: SqlAlarmRuleRespListSqlAlarmLevel{
			value: "Critical",
		},
	}
}

func (c SqlAlarmRuleRespListSqlAlarmLevel) Value() string {
	return c.value
}

func (c SqlAlarmRuleRespListSqlAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlAlarmRuleRespListSqlAlarmLevel) UnmarshalJSON(b []byte) error {
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

type SqlAlarmRuleRespListStatus struct {
	value string
}

type SqlAlarmRuleRespListStatusEnum struct {
	RUNNING  SqlAlarmRuleRespListStatus
	STOPPING SqlAlarmRuleRespListStatus
}

func GetSqlAlarmRuleRespListStatusEnum() SqlAlarmRuleRespListStatusEnum {
	return SqlAlarmRuleRespListStatusEnum{
		RUNNING: SqlAlarmRuleRespListStatus{
			value: "RUNNING 启用",
		},
		STOPPING: SqlAlarmRuleRespListStatus{
			value: "STOPPING 停止",
		},
	}
}

func (c SqlAlarmRuleRespListStatus) Value() string {
	return c.value
}

func (c SqlAlarmRuleRespListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlAlarmRuleRespListStatus) UnmarshalJSON(b []byte) error {
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

type SqlAlarmRuleRespListNotificationFrequency struct {
	value int32
}

type SqlAlarmRuleRespListNotificationFrequencyEnum struct {
	E_0   SqlAlarmRuleRespListNotificationFrequency
	E_5   SqlAlarmRuleRespListNotificationFrequency
	E_10  SqlAlarmRuleRespListNotificationFrequency
	E_15  SqlAlarmRuleRespListNotificationFrequency
	E_30  SqlAlarmRuleRespListNotificationFrequency
	E_60  SqlAlarmRuleRespListNotificationFrequency
	E_180 SqlAlarmRuleRespListNotificationFrequency
	E_360 SqlAlarmRuleRespListNotificationFrequency
}

func GetSqlAlarmRuleRespListNotificationFrequencyEnum() SqlAlarmRuleRespListNotificationFrequencyEnum {
	return SqlAlarmRuleRespListNotificationFrequencyEnum{
		E_0: SqlAlarmRuleRespListNotificationFrequency{
			value: 0,
		}, E_5: SqlAlarmRuleRespListNotificationFrequency{
			value: 5,
		}, E_10: SqlAlarmRuleRespListNotificationFrequency{
			value: 10,
		}, E_15: SqlAlarmRuleRespListNotificationFrequency{
			value: 15,
		}, E_30: SqlAlarmRuleRespListNotificationFrequency{
			value: 30,
		}, E_60: SqlAlarmRuleRespListNotificationFrequency{
			value: 60,
		}, E_180: SqlAlarmRuleRespListNotificationFrequency{
			value: 180,
		}, E_360: SqlAlarmRuleRespListNotificationFrequency{
			value: 360,
		},
	}
}

func (c SqlAlarmRuleRespListNotificationFrequency) Value() int32 {
	return c.value
}

func (c SqlAlarmRuleRespListNotificationFrequency) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlAlarmRuleRespListNotificationFrequency) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
