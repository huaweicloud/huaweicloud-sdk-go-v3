package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAlertRulesRequest Request Object
type ListAlertRulesRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`

	// offset
	Offset int64 `json:"offset"`

	// limit
	Limit int64 `json:"limit"`

	// sort_key
	SortKey *string `json:"sort_key,omitempty"`

	// sort_dir. asc, desc
	SortDir *ListAlertRulesRequestSortDir `json:"sort_dir,omitempty"`

	// pipe_id
	PipeId *string `json:"pipe_id,omitempty"`

	// rule_name
	RuleName *string `json:"rule_name,omitempty"`

	// rule_id
	RuleId *string `json:"rule_id,omitempty"`

	// status. ENABLED, DISABLED
	Status *[]ListAlertRulesRequestStatus `json:"status,omitempty"`

	// severity. TIPS, LOW, MEDIUM, HIGH, FATAL
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
