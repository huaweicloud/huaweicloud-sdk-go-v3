package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAlertRuleResponse Response Object
type UpdateAlertRuleResponse struct {

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
	QueryType *UpdateAlertRuleResponseQueryType `json:"query_type,omitempty"`

	// status. ENABLED, DISABLED
	Status *UpdateAlertRuleResponseStatus `json:"status,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *UpdateAlertRuleResponseSeverity `json:"severity,omitempty"`

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

func (o UpdateAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlertRuleResponse", string(data)}, " ")
}

type UpdateAlertRuleResponseQueryType struct {
	value string
}

type UpdateAlertRuleResponseQueryTypeEnum struct {
	SQL  UpdateAlertRuleResponseQueryType
	CBSL UpdateAlertRuleResponseQueryType
}

func GetUpdateAlertRuleResponseQueryTypeEnum() UpdateAlertRuleResponseQueryTypeEnum {
	return UpdateAlertRuleResponseQueryTypeEnum{
		SQL: UpdateAlertRuleResponseQueryType{
			value: "SQL",
		},
		CBSL: UpdateAlertRuleResponseQueryType{
			value: "CBSL",
		},
	}
}

func (c UpdateAlertRuleResponseQueryType) Value() string {
	return c.value
}

func (c UpdateAlertRuleResponseQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlertRuleResponseQueryType) UnmarshalJSON(b []byte) error {
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

type UpdateAlertRuleResponseStatus struct {
	value string
}

type UpdateAlertRuleResponseStatusEnum struct {
	ENABLED  UpdateAlertRuleResponseStatus
	DISABLED UpdateAlertRuleResponseStatus
}

func GetUpdateAlertRuleResponseStatusEnum() UpdateAlertRuleResponseStatusEnum {
	return UpdateAlertRuleResponseStatusEnum{
		ENABLED: UpdateAlertRuleResponseStatus{
			value: "ENABLED",
		},
		DISABLED: UpdateAlertRuleResponseStatus{
			value: "DISABLED",
		},
	}
}

func (c UpdateAlertRuleResponseStatus) Value() string {
	return c.value
}

func (c UpdateAlertRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlertRuleResponseStatus) UnmarshalJSON(b []byte) error {
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

type UpdateAlertRuleResponseSeverity struct {
	value string
}

type UpdateAlertRuleResponseSeverityEnum struct {
	TIPS   UpdateAlertRuleResponseSeverity
	LOW    UpdateAlertRuleResponseSeverity
	MEDIUM UpdateAlertRuleResponseSeverity
	HIGH   UpdateAlertRuleResponseSeverity
	FATAL  UpdateAlertRuleResponseSeverity
}

func GetUpdateAlertRuleResponseSeverityEnum() UpdateAlertRuleResponseSeverityEnum {
	return UpdateAlertRuleResponseSeverityEnum{
		TIPS: UpdateAlertRuleResponseSeverity{
			value: "TIPS",
		},
		LOW: UpdateAlertRuleResponseSeverity{
			value: "LOW",
		},
		MEDIUM: UpdateAlertRuleResponseSeverity{
			value: "MEDIUM",
		},
		HIGH: UpdateAlertRuleResponseSeverity{
			value: "HIGH",
		},
		FATAL: UpdateAlertRuleResponseSeverity{
			value: "FATAL",
		},
	}
}

func (c UpdateAlertRuleResponseSeverity) Value() string {
	return c.value
}

func (c UpdateAlertRuleResponseSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAlertRuleResponseSeverity) UnmarshalJSON(b []byte) error {
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
