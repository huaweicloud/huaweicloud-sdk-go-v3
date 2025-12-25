package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlertRulesRequest Request Object
type ListAlertRulesRequest struct {

	// 工作空间 ID
	WorkspaceId string `json:"workspace_id"`

	// 偏移量
	Offset *int64 `json:"offset,omitempty"`

	// 条数
	Limit *int64 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListAlertRulesRequestSortDir `json:"sort_dir,omitempty"`

	// 数据管道 ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 告警规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 告警规则 ID
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
	Status *[]ListAlertRulesRequestStatus `json:"status,omitempty"`

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Severity *[]ListAlertRulesRequestSeverity `json:"severity,omitempty"`
}

func (o ListAlertRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRulesRequest struct{}"
	}

	return strings.Join([]string{"ListAlertRulesRequest", string(data)}, " ")
}

type ListAlertRulesRequestSortDir struct {
	value string
}

type ListAlertRulesRequestSortDirEnum struct {
	ASC  ListAlertRulesRequestSortDir
	DESC ListAlertRulesRequestSortDir
}

func GetListAlertRulesRequestSortDirEnum() ListAlertRulesRequestSortDirEnum {
	return ListAlertRulesRequestSortDirEnum{
		ASC: ListAlertRulesRequestSortDir{
			value: "asc",
		},
		DESC: ListAlertRulesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListAlertRulesRequestSortDir) Value() string {
	return c.value
}

func (c ListAlertRulesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRulesRequestSortDir) UnmarshalJSON(b []byte) error {
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

type ListAlertRulesRequestStatus struct {
	value string
}

type ListAlertRulesRequestStatusEnum struct {
	ENABLED  ListAlertRulesRequestStatus
	DISABLED ListAlertRulesRequestStatus
}

func GetListAlertRulesRequestStatusEnum() ListAlertRulesRequestStatusEnum {
	return ListAlertRulesRequestStatusEnum{
		ENABLED: ListAlertRulesRequestStatus{
			value: "ENABLED",
		},
		DISABLED: ListAlertRulesRequestStatus{
			value: "DISABLED",
		},
	}
}

func (c ListAlertRulesRequestStatus) Value() string {
	return c.value
}

func (c ListAlertRulesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRulesRequestStatus) UnmarshalJSON(b []byte) error {
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

type ListAlertRulesRequestSeverity struct {
	value string
}

type ListAlertRulesRequestSeverityEnum struct {
	TIPS   ListAlertRulesRequestSeverity
	LOW    ListAlertRulesRequestSeverity
	MEDIUM ListAlertRulesRequestSeverity
	HIGH   ListAlertRulesRequestSeverity
	FATAL  ListAlertRulesRequestSeverity
}

func GetListAlertRulesRequestSeverityEnum() ListAlertRulesRequestSeverityEnum {
	return ListAlertRulesRequestSeverityEnum{
		TIPS: ListAlertRulesRequestSeverity{
			value: "TIPS",
		},
		LOW: ListAlertRulesRequestSeverity{
			value: "LOW",
		},
		MEDIUM: ListAlertRulesRequestSeverity{
			value: "MEDIUM",
		},
		HIGH: ListAlertRulesRequestSeverity{
			value: "HIGH",
		},
		FATAL: ListAlertRulesRequestSeverity{
			value: "FATAL",
		},
	}
}

func (c ListAlertRulesRequestSeverity) Value() string {
	return c.value
}

func (c ListAlertRulesRequestSeverity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRulesRequestSeverity) UnmarshalJSON(b []byte) error {
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
