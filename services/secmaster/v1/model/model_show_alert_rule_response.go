package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAlertRuleResponse Response Object
type ShowAlertRuleResponse struct {

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 自定义扩展信息
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 删除时间
	DeleteTime *int64 `json:"delete_time,omitempty"`

	// 告警分组
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 数据管道 ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 数据管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 查询语句
	Query *string `json:"query,omitempty"`

	// **参数解释**: 查询语法类型 - SQL查询 **约束限制**: 不涉及  **取值范围**: - SQL **默认值**:  SQL
	QueryType *ShowAlertRuleResponseQueryType `json:"query_type,omitempty"`

	// 告警规则 ID
	RuleId *string `json:"rule_id,omitempty"`

	// 告警规则名称
	RuleName *string `json:"rule_name,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// **参数解释**: 状态 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命 **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL **默认值** MEDIUM
	Severity *ShowAlertRuleResponseSeverity `json:"severity,omitempty"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status *ShowAlertRuleResponseStatus `json:"status,omitempty"`

	// 告警触发规则
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

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
	SQL ShowAlertRuleResponseQueryType
}

func GetShowAlertRuleResponseQueryTypeEnum() ShowAlertRuleResponseQueryTypeEnum {
	return ShowAlertRuleResponseQueryTypeEnum{
		SQL: ShowAlertRuleResponseQueryType{
			value: "SQL",
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
