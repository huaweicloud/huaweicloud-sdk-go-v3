package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAlertRuleRequestBody struct {

	// 数据管道 ID。Pipe ID.
	PipeId string `json:"pipe_id"`

	// 告警规则名称。Alert rule name.
	RuleName string `json:"rule_name"`

	// 描述。Description.
	Description *string `json:"description,omitempty"`

	// 查询语句。Query.
	Query string `json:"query"`

	// 查询语法，SQL。Query type. SQL.
	QueryType *CreateAlertRuleRequestBodyQueryType `json:"query_type,omitempty"`

	// 启用状态，启用、停用。Status, enabled, disabled.
	Status *CreateAlertRuleRequestBodyStatus `json:"status,omitempty"`

	// 严重程度，提示、低危、中危、高危、致命。Severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *CreateAlertRuleRequestBodySeverity `json:"severity,omitempty"`

	// 自定义扩展信息。Custom properties.
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 告警类型。Alert type.
	AlertType map[string]string `json:"alert_type,omitempty"`

	// 告警分组。Event grouping.
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 告警抑制。Suspression.
	Suspression *bool `json:"suspression,omitempty"`

	// 模拟告警。Simulation.
	Simulation *bool `json:"simulation,omitempty"`

	Schedule *Schedule `json:"schedule"`

	// 告警触发规则。Alert triggers.
	Triggers []AlertRuleTrigger `json:"triggers"`

	// 管道名称
	PipeName string `json:"pipe_name"`

	// 告警名称
	AlertName string `json:"alert_name"`

	// 告警描述
	AlertDescription *string `json:"alert_description,omitempty"`

	// 修复建议
	AlertRemediation *string `json:"alert_remediation,omitempty"`

	// 执行次数
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`
}

func (o CreateAlertRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAlertRuleRequestBody", string(data)}, " ")
}

type CreateAlertRuleRequestBodyQueryType struct {
	value string
}

type CreateAlertRuleRequestBodyQueryTypeEnum struct {
	SQL CreateAlertRuleRequestBodyQueryType
}

func GetCreateAlertRuleRequestBodyQueryTypeEnum() CreateAlertRuleRequestBodyQueryTypeEnum {
	return CreateAlertRuleRequestBodyQueryTypeEnum{
		SQL: CreateAlertRuleRequestBodyQueryType{
			value: "SQL",
		},
	}
}

func (c CreateAlertRuleRequestBodyQueryType) Value() string {
	return c.value
}

func (c CreateAlertRuleRequestBodyQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlertRuleRequestBodyQueryType) UnmarshalJSON(b []byte) error {
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

type CreateAlertRuleRequestBodyStatus struct {
	value string
}

type CreateAlertRuleRequestBodyStatusEnum struct {
	ENABLED  CreateAlertRuleRequestBodyStatus
	DISABLED CreateAlertRuleRequestBodyStatus
}

func GetCreateAlertRuleRequestBodyStatusEnum() CreateAlertRuleRequestBodyStatusEnum {
	return CreateAlertRuleRequestBodyStatusEnum{
		ENABLED: CreateAlertRuleRequestBodyStatus{
			value: "ENABLED",
		},
		DISABLED: CreateAlertRuleRequestBodyStatus{
			value: "DISABLED",
		},
	}
}

func (c CreateAlertRuleRequestBodyStatus) Value() string {
	return c.value
}

func (c CreateAlertRuleRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlertRuleRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

type CreateAlertRuleRequestBodySeverity struct {
	value string
}

type CreateAlertRuleRequestBodySeverityEnum struct {
	TIPS   CreateAlertRuleRequestBodySeverity
	LOW    CreateAlertRuleRequestBodySeverity
	MEDIUM CreateAlertRuleRequestBodySeverity
	HIGH   CreateAlertRuleRequestBodySeverity
	FATAL  CreateAlertRuleRequestBodySeverity
}

func GetCreateAlertRuleRequestBodySeverityEnum() CreateAlertRuleRequestBodySeverityEnum {
	return CreateAlertRuleRequestBodySeverityEnum{
		TIPS: CreateAlertRuleRequestBodySeverity{
			value: "TIPS",
		},
		LOW: CreateAlertRuleRequestBodySeverity{
			value: "LOW",
		},
		MEDIUM: CreateAlertRuleRequestBodySeverity{
			value: "MEDIUM",
		},
		HIGH: CreateAlertRuleRequestBodySeverity{
			value: "HIGH",
		},
		FATAL: CreateAlertRuleRequestBodySeverity{
			value: "FATAL",
		},
	}
}

func (c CreateAlertRuleRequestBodySeverity) Value() string {
	return c.value
}

func (c CreateAlertRuleRequestBodySeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlertRuleRequestBodySeverity) UnmarshalJSON(b []byte) error {
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
