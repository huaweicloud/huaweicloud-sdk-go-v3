package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPlaybookAuditLogsRequest Request Object
type ListPlaybookAuditLogsRequest struct {

	// **参数解释：** 内容类型 - application/json    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json  **默认取值：** application/json
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 分页查询参数，用于指定查询结果的起始位置，从0开始，limit 和 offset相加需要小于10000
	Offset int64 `json:"offset"`

	// 分页查询参数，用于指定一次查询最多的结果数，从1开始，limit 和 offset相加需要小于10000
	Limit int64 `json:"limit"`

	// **参数解释：** 排序字段 - start_time    开始时间 - end_time 结束时间  **约束限制：** 不涉及 **取值范围：** - start_time - end_time  **默认取值：** 无
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序
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
