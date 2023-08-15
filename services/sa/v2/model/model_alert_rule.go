package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlertRule struct {

	// rule_id
	RuleId string `json:"rule_id"`

	// pipe_id
	PipeId string `json:"pipe_id"`

	// pipe_name
	PipeName string `json:"pipe_name"`

	// create_by
	CreateBy string `json:"create_by"`

	// create_time
	CreateTime int64 `json:"create_time"`

	// update_by
	UpdateBy string `json:"update_by"`

	// update_time
	UpdateTime int64 `json:"update_time"`

	// delete_time
	DeleteTime *int64 `json:"delete_time,omitempty"`

	// rule_name
	RuleName string `json:"rule_name"`

	// query
	Query *string `json:"query,omitempty"`

	// query_type. SQL, CBSL.
	QueryType *AlertRuleQueryType `json:"query_type,omitempty"`

	// status. ENABLED, DISABLED
	Status AlertRuleStatus `json:"status"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity AlertRuleSeverity `json:"severity"`

	// accumulated_times
	AccumulatedTimes *int32 `json:"accumulated_times,omitempty"`

	// custom_properties
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// event_grouping
	EventGrouping *bool `json:"event_grouping,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// triggers
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`
}

func (o AlertRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRule struct{}"
	}

	return strings.Join([]string{"AlertRule", string(data)}, " ")
}

type AlertRuleQueryType struct {
	value string
}

type AlertRuleQueryTypeEnum struct {
	SQL  AlertRuleQueryType
	CBSL AlertRuleQueryType
}

func GetAlertRuleQueryTypeEnum() AlertRuleQueryTypeEnum {
	return AlertRuleQueryTypeEnum{
		SQL: AlertRuleQueryType{
			value: "SQL",
		},
		CBSL: AlertRuleQueryType{
			value: "CBSL",
		},
	}
}

func (c AlertRuleQueryType) Value() string {
	return c.value
}

func (c AlertRuleQueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleQueryType) UnmarshalJSON(b []byte) error {
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

type AlertRuleStatus struct {
	value string
}

type AlertRuleStatusEnum struct {
	ENABLED  AlertRuleStatus
	DISABLED AlertRuleStatus
}

func GetAlertRuleStatusEnum() AlertRuleStatusEnum {
	return AlertRuleStatusEnum{
		ENABLED: AlertRuleStatus{
			value: "ENABLED",
		},
		DISABLED: AlertRuleStatus{
			value: "DISABLED",
		},
	}
}

func (c AlertRuleStatus) Value() string {
	return c.value
}

func (c AlertRuleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleStatus) UnmarshalJSON(b []byte) error {
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

type AlertRuleSeverity struct {
	value string
}

type AlertRuleSeverityEnum struct {
	TIPS   AlertRuleSeverity
	LOW    AlertRuleSeverity
	MEDIUM AlertRuleSeverity
	HIGH   AlertRuleSeverity
	FATAL  AlertRuleSeverity
}

func GetAlertRuleSeverityEnum() AlertRuleSeverityEnum {
	return AlertRuleSeverityEnum{
		TIPS: AlertRuleSeverity{
			value: "TIPS",
		},
		LOW: AlertRuleSeverity{
			value: "LOW",
		},
		MEDIUM: AlertRuleSeverity{
			value: "MEDIUM",
		},
		HIGH: AlertRuleSeverity{
			value: "HIGH",
		},
		FATAL: AlertRuleSeverity{
			value: "FATAL",
		},
	}
}

func (c AlertRuleSeverity) Value() string {
	return c.value
}

func (c AlertRuleSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleSeverity) UnmarshalJSON(b []byte) error {
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
