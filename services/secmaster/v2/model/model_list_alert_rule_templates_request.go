package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlertRuleTemplatesRequest Request Object
type ListAlertRuleTemplatesRequest struct {

	// 工作空间 ID。Workspace ID.
	WorkspaceId string `json:"workspace_id"`

	// 偏移量。Offset.
	Offset int64 `json:"offset"`

	// 条数。Limit.
	Limit int64 `json:"limit"`

	// 排序字段。Sort key
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序，顺序、逆序。Sort direction, asc, desc。
	SortDir *ListAlertRuleTemplatesRequestSortDir `json:"sort_dir,omitempty"`

	// 严重程度，提示、低危、中危、高危、致命。Severity. TIPS, LOW, MEDIUM, HIGH, FATAL
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
