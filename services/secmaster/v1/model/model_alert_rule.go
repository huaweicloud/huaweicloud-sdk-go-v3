package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRule struct {

	// 创建人
	CreateBy string `json:"create_by"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 自定义扩展信息
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 删除时间
	DeleteTime *int64 `json:"delete_time,omitempty"`

	// 告警分组
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 数据管道 ID
	PipeId string `json:"pipe_id"`

	// 数据管道名称
	PipeName string `json:"pipe_name"`

	// 查询语句
	Query *string `json:"query,omitempty"`

	// **参数解释**: 查询语法类型 - SQL查询 **约束限制**: 不涉及  **取值范围**: - SQL **默认值**:  SQL
	QueryType *AlertRuleQueryType `json:"query_type,omitempty"`

	// 告警规则 ID
	RuleId string `json:"rule_id"`

	// 告警规则名称
	RuleName string `json:"rule_name"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// **参数解释**: 状态 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命 **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL **默认值** MEDIUM
	Severity AlertRuleSeverity `json:"severity"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status AlertRuleStatus `json:"status"`

	// 告警触发规则
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`

	// 更新人
	UpdateBy string `json:"update_by"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`
}

func (o AlertRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRule struct{}"
	}

	return strings.Join([]string{"AlertRule", string(data)}, " ")
}

type AlertRuleQueryType struct {
	value string
}

type AlertRuleQueryTypeEnum struct {
	SQL AlertRuleQueryType
}

func GetAlertRuleQueryTypeEnum() AlertRuleQueryTypeEnum {
	return AlertRuleQueryTypeEnum{
		SQL: AlertRuleQueryType{
			value: "SQL",
		},
	}
}

func (c AlertRuleQueryType) Value() string {
	return c.value
}

func (c AlertRuleQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleQueryType) UnmarshalJSON(b []byte) error {
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

type AlertRuleSeverity struct {
	value string
}

type AlertRuleSeverityEnum struct {
	TIPS   AlertRuleSeverity
	LOW    AlertRuleSeverity
	MEDIUM AlertRuleSeverity
	HIGH   AlertRuleSeverity
	FATAL  AlertRuleSeverity
}

func GetAlertRuleSeverityEnum() AlertRuleSeverityEnum {
	return AlertRuleSeverityEnum{
		TIPS: AlertRuleSeverity{
			value: "TIPS",
		},
		LOW: AlertRuleSeverity{
			value: "LOW",
		},
		MEDIUM: AlertRuleSeverity{
			value: "MEDIUM",
		},
		HIGH: AlertRuleSeverity{
			value: "HIGH",
		},
		FATAL: AlertRuleSeverity{
			value: "FATAL",
		},
	}
}

func (c AlertRuleSeverity) Value() string {
	return c.value
}

func (c AlertRuleSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleSeverity) UnmarshalJSON(b []byte) error {
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

type AlertRuleStatus struct {
	value string
}

type AlertRuleStatusEnum struct {
	ENABLED  AlertRuleStatus
	DISABLED AlertRuleStatus
}

func GetAlertRuleStatusEnum() AlertRuleStatusEnum {
	return AlertRuleStatusEnum{
		ENABLED: AlertRuleStatus{
			value: "ENABLED",
		},
		DISABLED: AlertRuleStatus{
			value: "DISABLED",
		},
	}
}

func (c AlertRuleStatus) Value() string {
	return c.value
}

func (c AlertRuleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleStatus) UnmarshalJSON(b []byte) error {
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
