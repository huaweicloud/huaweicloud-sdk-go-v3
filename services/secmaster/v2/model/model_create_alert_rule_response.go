package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAlertRuleResponse Response Object
type CreateAlertRuleResponse struct {

	// rule_id
	RuleId *string `json:"rule_id,omitempty"`

	// pipe_id
	PipeId *string `json:"pipe_id,omitempty"`

	// pipe_name
	PipeName *string `json:"pipe_name,omitempty"`

	// create_by
	CreateBy *string `json:"create_by,omitempty"`

	// create_time
	CreateTime *int64 `json:"create_time,omitempty"`

	// update_by
	UpdateBy *string `json:"update_by,omitempty"`

	// update_time
	UpdateTime *int64 `json:"update_time,omitempty"`

	// delete_time
	DeleteTime *int64 `json:"delete_time,omitempty"`

	// rule_name
	RuleName *string `json:"rule_name,omitempty"`

	// query
	Query *string `json:"query,omitempty"`

	// query_type. SQL, CBSL.
	QueryType *CreateAlertRuleResponseQueryType `json:"query_type,omitempty"`

	// status. ENABLED, DISABLED
	Status *CreateAlertRuleResponseStatus `json:"status,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *CreateAlertRuleResponseSeverity `json:"severity,omitempty"`

	// accumulated_times
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

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

func (o CreateAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAlertRuleResponse", string(data)}, " ")
}

type CreateAlertRuleResponseQueryType struct {
	value string
}

type CreateAlertRuleResponseQueryTypeEnum struct {
	SQL  CreateAlertRuleResponseQueryType
	CBSL CreateAlertRuleResponseQueryType
}

func GetCreateAlertRuleResponseQueryTypeEnum() CreateAlertRuleResponseQueryTypeEnum {
	return CreateAlertRuleResponseQueryTypeEnum{
		SQL: CreateAlertRuleResponseQueryType{
			value: "SQL",
		},
		CBSL: CreateAlertRuleResponseQueryType{
			value: "CBSL",
		},
	}
}

func (c CreateAlertRuleResponseQueryType) Value() string {
	return c.value
}

func (c CreateAlertRuleResponseQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlertRuleResponseQueryType) UnmarshalJSON(b []byte) error {
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

type CreateAlertRuleResponseStatus struct {
	value string
}

type CreateAlertRuleResponseStatusEnum struct {
	ENABLED  CreateAlertRuleResponseStatus
	DISABLED CreateAlertRuleResponseStatus
}

func GetCreateAlertRuleResponseStatusEnum() CreateAlertRuleResponseStatusEnum {
	return CreateAlertRuleResponseStatusEnum{
		ENABLED: CreateAlertRuleResponseStatus{
			value: "ENABLED",
		},
		DISABLED: CreateAlertRuleResponseStatus{
			value: "DISABLED",
		},
	}
}

func (c CreateAlertRuleResponseStatus) Value() string {
	return c.value
}

func (c CreateAlertRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlertRuleResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreateAlertRuleResponseSeverity struct {
	value string
}

type CreateAlertRuleResponseSeverityEnum struct {
	TIPS   CreateAlertRuleResponseSeverity
	LOW    CreateAlertRuleResponseSeverity
	MEDIUM CreateAlertRuleResponseSeverity
	HIGH   CreateAlertRuleResponseSeverity
	FATAL  CreateAlertRuleResponseSeverity
}

func GetCreateAlertRuleResponseSeverityEnum() CreateAlertRuleResponseSeverityEnum {
	return CreateAlertRuleResponseSeverityEnum{
		TIPS: CreateAlertRuleResponseSeverity{
			value: "TIPS",
		},
		LOW: CreateAlertRuleResponseSeverity{
			value: "LOW",
		},
		MEDIUM: CreateAlertRuleResponseSeverity{
			value: "MEDIUM",
		},
		HIGH: CreateAlertRuleResponseSeverity{
			value: "HIGH",
		},
		FATAL: CreateAlertRuleResponseSeverity{
			value: "FATAL",
		},
	}
}

func (c CreateAlertRuleResponseSeverity) Value() string {
	return c.value
}

func (c CreateAlertRuleResponseSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAlertRuleResponseSeverity) UnmarshalJSON(b []byte) error {
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
