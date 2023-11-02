package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAlertRuleResponse Response Object
type CreateAlertRuleResponse struct {

	// 告警规则 ID。Alert rule ID.
	RuleId *string `json:"rule_id,omitempty"`

	// 数据管道 ID。Pipe ID.
	PipeId *string `json:"pipe_id,omitempty"`

	// 数据管道名称。Pipe name.
	PipeName *string `json:"pipe_name,omitempty"`

	// 创建人。Create by.
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间。Create time.
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新人。Update by.
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间。Update time.
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 删除时间。Delete time.
	DeleteTime *int64 `json:"delete_time,omitempty"`

	// 告警规则名称。Alert rule name.
	RuleName *string `json:"rule_name,omitempty"`

	// 查询语句。Query.
	Query *string `json:"query,omitempty"`

	// 查询语法，SQL。Query type. SQL.
	QueryType *CreateAlertRuleResponseQueryType `json:"query_type,omitempty"`

	// 启用状态，启用、停用。Status, enabled, disabled.
	Status *CreateAlertRuleResponseStatus `json:"status,omitempty"`

	// 严重程度，提示、低危、中危、高危、致命。Severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *CreateAlertRuleResponseSeverity `json:"severity,omitempty"`

	// 自定义扩展信息。Custom properties.
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 告警分组。Event grouping.
	EventGrouping *bool `json:"event_grouping,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// 告警触发规则。Alert triggers.
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
	SQL CreateAlertRuleResponseQueryType
}

func GetCreateAlertRuleResponseQueryTypeEnum() CreateAlertRuleResponseQueryTypeEnum {
	return CreateAlertRuleResponseQueryTypeEnum{
		SQL: CreateAlertRuleResponseQueryType{
			value: "SQL",
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
