package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAlertRuleTemplateResponse Response Object
type ShowAlertRuleTemplateResponse struct {

	// 自定义扩展信息
	CustomProperties map[string]string `json:"custom_properties,omitempty"`

	// 数据源
	DataSource *string `json:"data_source,omitempty"`

	// 告警分组
	EventGrouping *bool `json:"event_grouping,omitempty"`

	// 查询语句
	Query *string `json:"query,omitempty"`

	// **参数解释**: 查询语法类型 - SQL查询 **约束限制**: 不涉及  **取值范围**: - SQL **默认值**:  SQL
	QueryType *ShowAlertRuleTemplateResponseQueryType `json:"query_type,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Severity *ShowAlertRuleTemplateResponseSeverity `json:"severity,omitempty"`

	// 告警规则模板 ID
	TemplateId *string `json:"template_id,omitempty"`

	// 告警规则模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 告警触发规则
	Triggers *[]AlertRuleTrigger `json:"triggers,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

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
	SQL ShowAlertRuleTemplateResponseQueryType
}

func GetShowAlertRuleTemplateResponseQueryTypeEnum() ShowAlertRuleTemplateResponseQueryTypeEnum {
	return ShowAlertRuleTemplateResponseQueryTypeEnum{
		SQL: ShowAlertRuleTemplateResponseQueryType{
			value: "SQL",
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
