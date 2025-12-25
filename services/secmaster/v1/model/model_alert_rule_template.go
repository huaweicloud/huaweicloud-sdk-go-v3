package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRuleTemplate struct {

	// 自定义扩展信息
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 数据源
	DataSource string `json:"data_source"`

	// 告警分组
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 查询语句
	Query *string `json:"query,omitempty"`

	// **参数解释**: 查询语法类型 - SQL查询 **约束限制**: 不涉及  **取值范围**: - SQL **默认值**:  SQL
	QueryType *AlertRuleTemplateQueryType `json:"query_type,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Severity AlertRuleTemplateSeverity `json:"severity"`

	// 告警规则模板 ID
	TemplateId string `json:"template_id"`

	// 告警规则模板名称
	TemplateName string `json:"template_name"`

	// 告警触发规则
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`

	// 更新时间
	UpdateTime int64 `json:"update_time"`

	// 版本
	Version string `json:"version"`
}

func (o AlertRuleTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleTemplate struct{}"
	}

	return strings.Join([]string{"AlertRuleTemplate", string(data)}, " ")
}

type AlertRuleTemplateQueryType struct {
	value string
}

type AlertRuleTemplateQueryTypeEnum struct {
	SQL AlertRuleTemplateQueryType
}

func GetAlertRuleTemplateQueryTypeEnum() AlertRuleTemplateQueryTypeEnum {
	return AlertRuleTemplateQueryTypeEnum{
		SQL: AlertRuleTemplateQueryType{
			value: "SQL",
		},
	}
}

func (c AlertRuleTemplateQueryType) Value() string {
	return c.value
}

func (c AlertRuleTemplateQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleTemplateQueryType) UnmarshalJSON(b []byte) error {
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

type AlertRuleTemplateSeverity struct {
	value string
}

type AlertRuleTemplateSeverityEnum struct {
	TIPS   AlertRuleTemplateSeverity
	LOW    AlertRuleTemplateSeverity
	MEDIUM AlertRuleTemplateSeverity
	HIGH   AlertRuleTemplateSeverity
	FATAL  AlertRuleTemplateSeverity
}

func GetAlertRuleTemplateSeverityEnum() AlertRuleTemplateSeverityEnum {
	return AlertRuleTemplateSeverityEnum{
		TIPS: AlertRuleTemplateSeverity{
			value: "TIPS",
		},
		LOW: AlertRuleTemplateSeverity{
			value: "LOW",
		},
		MEDIUM: AlertRuleTemplateSeverity{
			value: "MEDIUM",
		},
		HIGH: AlertRuleTemplateSeverity{
			value: "HIGH",
		},
		FATAL: AlertRuleTemplateSeverity{
			value: "FATAL",
		},
	}
}

func (c AlertRuleTemplateSeverity) Value() string {
	return c.value
}

func (c AlertRuleTemplateSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleTemplateSeverity) UnmarshalJSON(b []byte) error {
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
