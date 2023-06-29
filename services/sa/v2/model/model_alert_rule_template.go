package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRuleTemplate struct {

	// template_id
	TemplateId string `json:"template_id"`

	// update_time
	UpdateTime int64 `json:"update_time"`

	// template_name
	TemplateName string `json:"template_name"`

	// data_source
	DataSource string `json:"data_source"`

	// version
	Version string `json:"version"`

	// query
	Query *string `json:"query,omitempty"`

	// query_type. SQL, CBSL.
	QueryType *AlertRuleTemplateQueryType `json:"query_type,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity AlertRuleTemplateSeverity `json:"severity"`

	// custom_properties
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// event_grouping
	EventGrouping *bool `json:"event_grouping,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// triggers
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
	SQL  AlertRuleTemplateQueryType
	CBSL AlertRuleTemplateQueryType
}

func GetAlertRuleTemplateQueryTypeEnum() AlertRuleTemplateQueryTypeEnum {
	return AlertRuleTemplateQueryTypeEnum{
		SQL: AlertRuleTemplateQueryType{
			value: "SQL",
		},
		CBSL: AlertRuleTemplateQueryType{
			value: "CBSL",
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
