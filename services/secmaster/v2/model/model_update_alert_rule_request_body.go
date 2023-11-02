package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAlertRuleRequestBody struct {

	// 告警规则名称。Alert rule name.
	RuleName *string `json:"rule_name,omitempty"`

	// 描述。Description.
	Description *string `json:"description,omitempty"`

	// 查询语句。Query.
	Query *string `json:"query,omitempty"`

	// 查询语法，SQL。Query type. SQL.
	QueryType *UpdateAlertRuleRequestBodyQueryType `json:"query_type,omitempty"`

	// 启用状态，启用、停用。Status, enabled, disabled.
	Status *UpdateAlertRuleRequestBodyStatus `json:"status,omitempty"`

	// 严重程度，提示、低危、中危、高危、致命。Severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *UpdateAlertRuleRequestBodySeverity `json:"severity,omitempty"`

	// 自定义扩展信息。Custom properties.
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 告警类型。Alert type.
	AlertType map[string]string `json:"alert_type,omitempty"`

	// 告警分组。Event grouping.
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 告警抑制。Suppression
	Suppression *bool `json:"suppression,omitempty"`

	// 模拟告警。Simulation.
	Simulation *bool `json:"simulation,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// 告警触发规则。Alert triggers.
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`
}

func (o UpdateAlertRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlertRuleRequestBody", string(data)}, " ")
}

type UpdateAlertRuleRequestBodyQueryType struct {
	value string
}

type UpdateAlertRuleRequestBodyQueryTypeEnum struct {
	SQL UpdateAlertRuleRequestBodyQueryType
}

func GetUpdateAlertRuleRequestBodyQueryTypeEnum() UpdateAlertRuleRequestBodyQueryTypeEnum {
	return UpdateAlertRuleRequestBodyQueryTypeEnum{
		SQL: UpdateAlertRuleRequestBodyQueryType{
			value: "SQL",
		},
	}
}

func (c UpdateAlertRuleRequestBodyQueryType) Value() string {
	return c.value
}

func (c UpdateAlertRuleRequestBodyQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlertRuleRequestBodyQueryType) UnmarshalJSON(b []byte) error {
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

type UpdateAlertRuleRequestBodyStatus struct {
	value string
}

type UpdateAlertRuleRequestBodyStatusEnum struct {
	ENABLED  UpdateAlertRuleRequestBodyStatus
	DISABLED UpdateAlertRuleRequestBodyStatus
}

func GetUpdateAlertRuleRequestBodyStatusEnum() UpdateAlertRuleRequestBodyStatusEnum {
	return UpdateAlertRuleRequestBodyStatusEnum{
		ENABLED: UpdateAlertRuleRequestBodyStatus{
			value: "ENABLED",
		},
		DISABLED: UpdateAlertRuleRequestBodyStatus{
			value: "DISABLED",
		},
	}
}

func (c UpdateAlertRuleRequestBodyStatus) Value() string {
	return c.value
}

func (c UpdateAlertRuleRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlertRuleRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

type UpdateAlertRuleRequestBodySeverity struct {
	value string
}

type UpdateAlertRuleRequestBodySeverityEnum struct {
	TIPS   UpdateAlertRuleRequestBodySeverity
	LOW    UpdateAlertRuleRequestBodySeverity
	MEDIUM UpdateAlertRuleRequestBodySeverity
	HIGH   UpdateAlertRuleRequestBodySeverity
	FATAL  UpdateAlertRuleRequestBodySeverity
}

func GetUpdateAlertRuleRequestBodySeverityEnum() UpdateAlertRuleRequestBodySeverityEnum {
	return UpdateAlertRuleRequestBodySeverityEnum{
		TIPS: UpdateAlertRuleRequestBodySeverity{
			value: "TIPS",
		},
		LOW: UpdateAlertRuleRequestBodySeverity{
			value: "LOW",
		},
		MEDIUM: UpdateAlertRuleRequestBodySeverity{
			value: "MEDIUM",
		},
		HIGH: UpdateAlertRuleRequestBodySeverity{
			value: "HIGH",
		},
		FATAL: UpdateAlertRuleRequestBodySeverity{
			value: "FATAL",
		},
	}
}

func (c UpdateAlertRuleRequestBodySeverity) Value() string {
	return c.value
}

func (c UpdateAlertRuleRequestBodySeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlertRuleRequestBodySeverity) UnmarshalJSON(b []byte) error {
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
