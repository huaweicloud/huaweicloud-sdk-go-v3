package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndexUsageExportTasksRequest Request Object
type ListIndexUsageExportTasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 当前页
	CurPage *int32 `json:"cur_page,omitempty"`

	// 页大小
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListIndexUsageExportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndexUsageExportTasksRequest struct{}"
	}

	return strings.Join([]string{"ListIndexUsageExportTasksRequest", string(data)}, " ")
}
