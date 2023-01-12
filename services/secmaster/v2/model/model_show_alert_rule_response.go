package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowAlertRuleResponse struct {

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
	QueryType *ShowAlertRuleResponseQueryType `json:"query_type,omitempty"`

	// status. ENABLED, DISABLED
	Status *ShowAlertRuleResponseStatus `json:"status,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *ShowAlertRuleResponseSeverity `json:"severity,omitempty"`

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

func (o ShowAlertRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowAlertRuleResponse", string(data)}, " ")
}

type ShowAlertRuleResponseQueryType struct {
	value string
}

type ShowAlertRuleResponseQueryTypeEnum struct {
	SQL  ShowAlertRuleResponseQueryType
	CBSL ShowAlertRuleResponseQueryType
}

func GetShowAlertRuleResponseQueryTypeEnum() ShowAlertRuleResponseQueryTypeEnum {
	return ShowAlertRuleResponseQueryTypeEnum{
		SQL: ShowAlertRuleResponseQueryType{
			value: "SQL",
		},
		CBSL: ShowAlertRuleResponseQueryType{
			value: "CBSL",
		},
	}
}

func (c ShowAlertRuleResponseQueryType) Value() string {
	return c.value
}

func (c ShowAlertRuleResponseQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlertRuleResponseQueryType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowAlertRuleResponseStatus struct {
	value string
}

type ShowAlertRuleResponseStatusEnum struct {
	ENABLED  ShowAlertRuleResponseStatus
	DISABLED ShowAlertRuleResponseStatus
}

func GetShowAlertRuleResponseStatusEnum() ShowAlertRuleResponseStatusEnum {
	return ShowAlertRuleResponseStatusEnum{
		ENABLED: ShowAlertRuleResponseStatus{
			value: "ENABLED",
		},
		DISABLED: ShowAlertRuleResponseStatus{
			value: "DISABLED",
		},
	}
}

func (c ShowAlertRuleResponseStatus) Value() string {
	return c.value
}

func (c ShowAlertRuleResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlertRuleResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowAlertRuleResponseSeverity struct {
	value string
}

type ShowAlertRuleResponseSeverityEnum struct {
	TIPS   ShowAlertRuleResponseSeverity
	LOW    ShowAlertRuleResponseSeverity
	MEDIUM ShowAlertRuleResponseSeverity
	HIGH   ShowAlertRuleResponseSeverity
	FATAL  ShowAlertRuleResponseSeverity
}

func GetShowAlertRuleResponseSeverityEnum() ShowAlertRuleResponseSeverityEnum {
	return ShowAlertRuleResponseSeverityEnum{
		TIPS: ShowAlertRuleResponseSeverity{
			value: "TIPS",
		},
		LOW: ShowAlertRuleResponseSeverity{
			value: "LOW",
		},
		MEDIUM: ShowAlertRuleResponseSeverity{
			value: "MEDIUM",
		},
		HIGH: ShowAlertRuleResponseSeverity{
			value: "HIGH",
		},
		FATAL: ShowAlertRuleResponseSeverity{
			value: "FATAL",
		},
	}
}

func (c ShowAlertRuleResponseSeverity) Value() string {
	return c.value
}

func (c ShowAlertRuleResponseSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAlertRuleResponseSeverity) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
