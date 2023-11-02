package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRuleTemplate struct {

	// 告警规则模板 ID。Alert rule template ID.
	TemplateId string `json:"template_id"`

	// 更新时间。Update time.
	UpdateTime int64 `json:"update_time"`

	// 告警规则模板名称。Alert rule template name.
	TemplateName string `json:"template_name"`

	// 数据源。Data source.
	DataSource string `json:"data_source"`

	// 版本。Version
	Version string `json:"version"`

	// 查询语句。Query.
	Query *string `json:"query,omitempty"`

	// 查询语法，SQL。Query type. SQL.
	QueryType *AlertRuleTemplateQueryType `json:"query_type,omitempty"`

	// 严重程度，提示、低危、中危、高危、致命。Severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity AlertRuleTemplateSeverity `json:"severity"`

	// 自定义扩展信息。Custom properties.
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 告警分组。Event grouping.
	EventGrouping *bool `json:"event_grouping,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// 告警触发规则。Alert triggers.
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`
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
