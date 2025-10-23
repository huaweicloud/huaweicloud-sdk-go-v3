package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskStatusSummaryResponseBodyFilter 过滤条件
type ShowTaskStatusSummaryResponseBodyFilter struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	TaskType *TaskTypeEnum `json:"task_type,omitempty"`

	ResourceType *ResourceTypeEnum `json:"resource_type,omitempty"`

	// 起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowTaskStatusSummaryResponseBodyFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskStatusSummaryResponseBodyFilter struct{}"
	}

	return strings.Join([]string{"ShowTaskStatusSummaryResponseBodyFilter", string(data)}, " ")
}
