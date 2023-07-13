package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPlaybookAuditLogsRequest Request Object
type ListPlaybookAuditLogsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// offset
	Offset int64 `json:"offset"`

	// limit
	Limit int64 `json:"limit"`

	// sort_key
	SortKey *string `json:"sort_key,omitempty"`

	// sort_dir. asc, desc
	SortDir *ListPlaybookAuditLogsRequestSortDir `json:"sort_dir,omitempty"`

	Body *AuditLogInfo `json:"body,omitempty"`
}

func (o ListPlaybookAuditLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookAuditLogsRequest struct{}"
	}

	return strings.Join([]string{"ListPlaybookAuditLogsRequest", string(data)}, " ")
}

type ListPlaybookAuditLogsRequestSortDir struct {
	value string
}

type ListPlaybookAuditLogsRequestSortDirEnum struct {
	ASC  ListPlaybookAuditLogsRequestSortDir
	DESC ListPlaybookAuditLogsRequestSortDir
}

func GetListPlaybookAuditLogsRequestSortDirEnum() ListPlaybookAuditLogsRequestSortDirEnum {
	return ListPlaybookAuditLogsRequestSortDirEnum{
		ASC: ListPlaybookAuditLogsRequestSortDir{
			value: "asc",
		},
		DESC: ListPlaybookAuditLogsRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListPlaybookAuditLogsRequestSortDir) Value() string {
	return c.value
}

func (c ListPlaybookAuditLogsRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPlaybookAuditLogsRequestSortDir) UnmarshalJSON(b []byte) error {
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
