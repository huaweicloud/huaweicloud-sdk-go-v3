package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAlertRuleRequestBody struct {

	// 执行次数
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

	// 告警描述
	AlertDescription *string `json:"alert_description,omitempty"`

	// 告警名称
	AlertName string `json:"alert_name"`

	// 修复建议
	AlertRemediation *string `json:"alert_remediation,omitempty"`

	// 告警类型，通过“查询数据类列表” 接口获取.
	AlertType map[string]string `json:"alert_type,omitempty"`

	// 自定义扩展信息
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 告警分组
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 数据管道 ID
	PipeId string `json:"pipe_id"`

	// 管道名称
	PipeName string `json:"pipe_name"`

	// 查询语句
	Query string `json:"query"`

	// **参数解释**: 查询语法类型 - SQL查询 **约束限制**: 不涉及  **取值范围**: - SQL **默认值**:  SQL
	QueryType *CreateAlertRuleRequestBodyQueryType `json:"query_type,omitempty"`

	// 告警规则名称
	RuleName string `json:"rule_name"`

	Schedule *Schedule `json:"schedule"`

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Severity *CreateAlertRuleRequestBodySeverity `json:"severity,omitempty"`

	// 模拟告警
	Simulation *bool `json:"simulation,omitempty"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status *CreateAlertRuleRequestBodyStatus `json:"status,omitempty"`

	// 告警抑制
	Suppression *bool `json:"suppression,omitempty"`

	// 告警触发规则
	Triggers []AlertRuleTrigger `json:"triggers"`
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
