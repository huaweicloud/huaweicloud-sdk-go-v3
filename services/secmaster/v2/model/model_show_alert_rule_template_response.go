package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAlertRuleTemplateResponse Response Object
type ShowAlertRuleTemplateResponse struct {

	// template_id
	TemplateId *string `json:"template_id,omitempty"`

	// update_time
	UpdateTime *int64 `json:"update_time,omitempty"`

	// template_name
	TemplateName *string `json:"template_name,omitempty"`

	// data_source
	DataSource *string `json:"data_source,omitempty"`

	// version
	Version *string `json:"version,omitempty"`

	// query
	Query *string `json:"query,omitempty"`

	// query_type. SQL, CBSL.
	QueryType *ShowAlertRuleTemplateResponseQueryType `json:"query_type,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *ShowAlertRuleTemplateResponseSeverity `json:"severity,omitempty"`

	// custom_properties
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// event_grouping
	EventGrouping *bool `json:"event_grouping,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// triggers
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAlertRuleTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRuleTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowAlertRuleTemplateResponse", string(data)}, " ")
}

type ShowAlertRuleTemplateResponseQueryType struct {
	value string
}

type ShowAlertRuleTemplateResponseQueryTypeEnum struct {
	SQL  ShowAlertRuleTemplateResponseQueryType
	CBSL ShowAlertRuleTemplateResponseQueryType
}

func GetShowAlertRuleTemplateResponseQueryTypeEnum() ShowAlertRuleTemplateResponseQueryTypeEnum {
	return ShowAlertRuleTemplateResponseQueryTypeEnum{
		SQL: ShowAlertRuleTemplateResponseQueryType{
			value: "SQL",
		},
		CBSL: ShowAlertRuleTemplateResponseQueryType{
			value: "CBSL",
		},
	}
}

func (c ShowAlertRuleTemplateResponseQueryType) Value() string {
	return c.value
}

func (c ShowAlertRuleTemplateResponseQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlertRuleTemplateResponseQueryType) UnmarshalJSON(b []byte) error {
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

type ShowAlertRuleTemplateResponseSeverity struct {
	value string
}

type ShowAlertRuleTemplateResponseSeverityEnum struct {
	TIPS   ShowAlertRuleTemplateResponseSeverity
	LOW    ShowAlertRuleTemplateResponseSeverity
	MEDIUM ShowAlertRuleTemplateResponseSeverity
	HIGH   ShowAlertRuleTemplateResponseSeverity
	FATAL  ShowAlertRuleTemplateResponseSeverity
}

func GetShowAlertRuleTemplateResponseSeverityEnum() ShowAlertRuleTemplateResponseSeverityEnum {
	return ShowAlertRuleTemplateResponseSeverityEnum{
		TIPS: ShowAlertRuleTemplateResponseSeverity{
			value: "TIPS",
		},
		LOW: ShowAlertRuleTemplateResponseSeverity{
			value: "LOW",
		},
		MEDIUM: ShowAlertRuleTemplateResponseSeverity{
			value: "MEDIUM",
		},
		HIGH: ShowAlertRuleTemplateResponseSeverity{
			value: "HIGH",
		},
		FATAL: ShowAlertRuleTemplateResponseSeverity{
			value: "FATAL",
		},
	}
}

func (c ShowAlertRuleTemplateResponseSeverity) Value() string {
	return c.value
}

func (c ShowAlertRuleTemplateResponseSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlertRuleTemplateResponseSeverity) UnmarshalJSON(b []byte) error {
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
