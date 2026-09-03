package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceHealthReportTasksRequest Request Object
type ListInstanceHealthReportTasksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 开始时间（Unix时间戳，毫秒）
	StartAt *string `json:"start_at,omitempty"`

	// 结束时间（Unix时间戳，毫秒）
	EndAt *string `json:"end_at,omitempty"`

	// 页码
	PageNum *string `json:"page_num,omitempty"`

	// 每页记录数
	PageSize *string `json:"page_size,omitempty"`
}

func (o ListInstanceHealthReportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceHealthReportTasksRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceHealthReportTasksRequest", string(data)}, " ")
}
