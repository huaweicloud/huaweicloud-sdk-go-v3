package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRule struct {

	// 告警规则 ID。Alert rule ID.
	RuleId string `json:"rule_id"`

	// 数据管道 ID。Pipe ID.
	PipeId string `json:"pipe_id"`

	// 数据管道名称。Pipe name.
	PipeName string `json:"pipe_name"`

	// 创建人。Create by.
	CreateBy string `json:"create_by"`

	// 创建时间。Create time.
	CreateTime int64 `json:"create_time"`

	// 更新人。Update by.
	UpdateBy string `json:"update_by"`

	// 更新时间。Update time.
	UpdateTime int64 `json:"update_time"`

	// 删除时间。Delete time.
	DeleteTime *int64 `json:"delete_time,omitempty"`

	// 告警规则名称。Alert rule name.
	RuleName string `json:"rule_name"`

	// 查询语句。Query.
	Query *string `json:"query,omitempty"`

	// 查询语法，SQL。Query type. SQL.
	QueryType *AlertRuleQueryType `json:"query_type,omitempty"`

	// 启用状态，启用、停用。Status, enabled, disabled.
	Status AlertRuleStatus `json:"status"`

	// 严重程度，提示、低危、中危、高危、致命。Severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity AlertRuleSeverity `json:"severity"`

	// 自定义扩展信息。Custom properties.
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 告警分组。Event grouping.
	EventGrouping *bool `json:"event_grouping,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// 告警触发规则。Alert triggers.
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`
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
