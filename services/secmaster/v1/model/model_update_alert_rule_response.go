package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAlertRuleResponse Response Object
type UpdateAlertRuleResponse struct {

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
	QueryType *UpdateAlertRuleResponseQueryType `json:"query_type,omitempty"`

	// 告警规则 ID
	RuleId *string `json:"rule_id,omitempty"`

	// 告警规则名称
	RuleName *string `json:"rule_name,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// **参数解释**: 状态 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命 **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL **默认值** MEDIUM
	Severity *UpdateAlertRuleResponseSeverity `json:"severity,omitempty"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status *UpdateAlertRuleResponseStatus `json:"status,omitempty"`

	// 告警触发规则
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

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
	SQL UpdateAlertRuleResponseQueryType
}

func GetUpdateAlertRuleResponseQueryTypeEnum() UpdateAlertRuleResponseQueryTypeEnum {
	return UpdateAlertRuleResponseQueryTypeEnum{
		SQL: UpdateAlertRuleResponseQueryType{
			value: "SQL",
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
