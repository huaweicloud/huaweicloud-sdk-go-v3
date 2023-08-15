package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAlertRuleRequestBody struct {

	// rule_name
	RuleName *string `json:"rule_name,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// query
	Query *string `json:"query,omitempty"`

	// query_type. SQL, CBSL.
	QueryType *UpdateAlertRuleRequestBodyQueryType `json:"query_type,omitempty"`

	// status. ENABLED, DISABLED
	Status *UpdateAlertRuleRequestBodyStatus `json:"status,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *UpdateAlertRuleRequestBodySeverity `json:"severity,omitempty"`

	// accumulated_times
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

	// custom_properties
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// alert_type
	AlertType map[string]string `json:"alert_type,omitempty"`

	// event_grouping
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// suppression
	Suppression *bool `json:"suppression,omitempty"`

	// simulation
	Simulation *bool `json:"simulation,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// triggers
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
	SQL  UpdateAlertRuleRequestBodyQueryType
	CBSL UpdateAlertRuleRequestBodyQueryType
}

func GetUpdateAlertRuleRequestBodyQueryTypeEnum() UpdateAlertRuleRequestBodyQueryTypeEnum {
	return UpdateAlertRuleRequestBodyQueryTypeEnum{
		SQL: UpdateAlertRuleRequestBodyQueryType{
			value: "SQL",
		},
		CBSL: UpdateAlertRuleRequestBodyQueryType{
			value: "CBSL",
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
