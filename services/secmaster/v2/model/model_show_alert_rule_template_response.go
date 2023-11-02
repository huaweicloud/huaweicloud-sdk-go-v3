package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAlertRuleTemplateResponse Response Object
type ShowAlertRuleTemplateResponse struct {

	// 告警规则模板 ID。Alert rule template ID.
	TemplateId *string `json:"template_id,omitempty"`

	// 更新时间。Update time.
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 告警规则模板名称。Alert rule template name.
	TemplateName *string `json:"template_name,omitempty"`

	// 数据源。Data source.
	DataSource *string `json:"data_source,omitempty"`

	// 版本。Version
	Version *string `json:"version,omitempty"`

	// 查询语句。Query.
	Query *string `json:"query,omitempty"`

	// 查询语法，SQL。Query type. SQL.
	QueryType *ShowAlertRuleTemplateResponseQueryType `json:"query_type,omitempty"`

	// 严重程度，提示、低危、中危、高危、致命。Severity. TIPS, LOW, MEDIUM, HIGH, FATAL
	Severity *ShowAlertRuleTemplateResponseSeverity `json:"severity,omitempty"`

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
