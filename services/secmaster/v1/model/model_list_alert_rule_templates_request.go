package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlertRuleTemplatesRequest Request Object
type ListAlertRuleTemplatesRequest struct {

	// 工作空间 ID
	WorkspaceId string `json:"workspace_id"`

	// 偏移量
	Offset int64 `json:"offset"`

	// 条数
	Limit int64 `json:"limit"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 目录排序 - asc 升序排列 - desc 降序排列  **约束限制** 不涉及 **取值范围**: - asc - desc  **默认值** 不涉及
	SortDir *ListAlertRuleTemplatesRequestSortDir `json:"sort_dir,omitempty"`

	// **参数解释**: 告警等级 - TIPS 提示 - LOW 低危 - MEDIUM 中危 - HIGH 高危 - FATAL 致命  **约束限制** 不涉及 **取值范围**: - TIPS - LOW - MEDIUM - HIGH - FATAL  **默认值** 不涉及
	Severity *[]ListAlertRuleTemplatesRequestSeverity `json:"severity,omitempty"`
}

func (o ListAlertRuleTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplatesRequest", string(data)}, " ")
}

type ListAlertRuleTemplatesRequestSortDir struct {
	value string
}

type ListAlertRuleTemplatesRequestSortDirEnum struct {
	ASC  ListAlertRuleTemplatesRequestSortDir
	DESC ListAlertRuleTemplatesRequestSortDir
}

func GetListAlertRuleTemplatesRequestSortDirEnum() ListAlertRuleTemplatesRequestSortDirEnum {
	return ListAlertRuleTemplatesRequestSortDirEnum{
		ASC: ListAlertRuleTemplatesRequestSortDir{
			value: "asc",
		},
		DESC: ListAlertRuleTemplatesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListAlertRuleTemplatesRequestSortDir) Value() string {
	return c.value
}

func (c ListAlertRuleTemplatesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAlertRuleTemplatesRequestSortDir) UnmarshalJSON(b []byte) error {
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
