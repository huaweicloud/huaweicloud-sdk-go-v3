package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlertRuleTemplatesRequest Request Object
type ListAlertRuleTemplatesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status *ListAlertRuleTemplatesRequestStatus `json:"status,omitempty"`

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Severity *ListAlertRuleTemplatesRequestSeverity `json:"severity,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int64 `json:"offset,omitempty"`

	// **参数解释：** 查询数据限制 **取值范围：** 0-1000 **默认取值：** 不涉及
	Limit *int64 `json:"limit,omitempty"`

	// 按照属性排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序，支持 `ASC` 或 `DESC`。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListAlertRuleTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplatesRequest", string(data)}, " ")
}

type ListAlertRuleTemplatesRequestStatus struct {
	value string
}

type ListAlertRuleTemplatesRequestStatusEnum struct {
	ENABLED  ListAlertRuleTemplatesRequestStatus
	DISABLED ListAlertRuleTemplatesRequestStatus
}

func GetListAlertRuleTemplatesRequestStatusEnum() ListAlertRuleTemplatesRequestStatusEnum {
	return ListAlertRuleTemplatesRequestStatusEnum{
		ENABLED: ListAlertRuleTemplatesRequestStatus{
			value: "ENABLED",
		},
		DISABLED: ListAlertRuleTemplatesRequestStatus{
			value: "DISABLED",
		},
	}
}

func (c ListAlertRuleTemplatesRequestStatus) Value() string {
	return c.value
}

func (c ListAlertRuleTemplatesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRuleTemplatesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListAlertRuleTemplatesRequestSeverity struct {
	value string
}

type ListAlertRuleTemplatesRequestSeverityEnum struct {
	TIPS   ListAlertRuleTemplatesRequestSeverity
	LOW    ListAlertRuleTemplatesRequestSeverity
	MEDIUM ListAlertRuleTemplatesRequestSeverity
	HIGH   ListAlertRuleTemplatesRequestSeverity
	FATAL  ListAlertRuleTemplatesRequestSeverity
}

func GetListAlertRuleTemplatesRequestSeverityEnum() ListAlertRuleTemplatesRequestSeverityEnum {
	return ListAlertRuleTemplatesRequestSeverityEnum{
		TIPS: ListAlertRuleTemplatesRequestSeverity{
			value: "TIPS",
		},
		LOW: ListAlertRuleTemplatesRequestSeverity{
			value: "LOW",
		},
		MEDIUM: ListAlertRuleTemplatesRequestSeverity{
			value: "MEDIUM",
		},
		HIGH: ListAlertRuleTemplatesRequestSeverity{
			value: "HIGH",
		},
		FATAL: ListAlertRuleTemplatesRequestSeverity{
			value: "FATAL",
		},
	}
}

func (c ListAlertRuleTemplatesRequestSeverity) Value() string {
	return c.value
}

func (c ListAlertRuleTemplatesRequestSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRuleTemplatesRequestSeverity) UnmarshalJSON(b []byte) error {
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
