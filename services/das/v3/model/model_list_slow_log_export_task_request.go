package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogExportTaskRequest Request Object
type ListSlowLogExportTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`

	// 导出类型，取值范围：slowsql、slowsqldetails
	ExportType *string `json:"export_type,omitempty"`
}

func (o ListSlowLogExportTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogExportTaskRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogExportTaskRequest", string(data)}, " ")
}
