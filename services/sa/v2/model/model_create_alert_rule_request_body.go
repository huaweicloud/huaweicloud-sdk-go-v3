package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAlertRuleRequestBody struct {

	// pipe_id
	PipeId string `json:"pipe_id"`

	// rule_name
	RuleName string `json:"rule_name"`

	// description
	Description *string `json:"description,omitempty"`

	// query
	Query string `json:"query"`

	// query_type. SQL, CBSL.
	QueryType *CreateAlertRuleRequestBodyQueryType `json:"query_type,omitempty"`

	// status. ENABLED, DISABLED
	Status *CreateAlertRuleRequestBodyStatus `json:"status,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *CreateAlertRuleRequestBodySeverity `json:"severity,omitempty"`

	// accumulated_times
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

	// custom_properties
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// alert_type
	AlertType map[string]string `json:"alert_type,omitempty"`

	// event_grouping
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// suspression
	Suspression *bool `json:"suspression,omitempty"`

	// simulation
	Simulation *bool `json:"simulation,omitempty"`

	Schedule *Schedule `json:"schedule"`

	// triggers
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
	SQL  CreateAlertRuleRequestBodyQueryType
	CBSL CreateAlertRuleRequestBodyQueryType
}

func GetCreateAlertRuleRequestBodyQueryTypeEnum() CreateAlertRuleRequestBodyQueryTypeEnum {
	return CreateAlertRuleRequestBodyQueryTypeEnum{
		SQL: CreateAlertRuleRequestBodyQueryType{
			value: "SQL",
		},
		CBSL: CreateAlertRuleRequestBodyQueryType{
			value: "CBSL",
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
