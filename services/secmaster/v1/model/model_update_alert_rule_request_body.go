package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAlertRuleRequestBody struct {

	// 告警类型，通过“查询数据类列表” 接口获取
	AlertType map[string]string `json:"alert_type,omitempty"`

	// 自定义扩展信息
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 告警分组
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 查询语句
	Query *string `json:"query,omitempty"`

	// **参数解释**: 查询语法类型 - SQL查询 **约束限制**: 不涉及  **取值范围**: - SQL **默认值**:  SQL
	QueryType *UpdateAlertRuleRequestBodyQueryType `json:"query_type,omitempty"`

	// 告警规则名称
	RuleName *string `json:"rule_name,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Severity *UpdateAlertRuleRequestBodySeverity `json:"severity,omitempty"`

	// 模拟告警
	Simulation *bool `json:"simulation,omitempty"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status *UpdateAlertRuleRequestBodyStatus `json:"status,omitempty"`

	// 告警抑制
	Suppression *bool `json:"suppression,omitempty"`

	// 告警触发规则
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
