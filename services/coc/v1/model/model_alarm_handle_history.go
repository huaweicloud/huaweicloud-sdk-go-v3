package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmHandleHistory 告警工单执行结果
type AlarmHandleHistory struct {

	// 执行工单id
	WorkOrderId *string `json:"work_order_id,omitempty"`

	// 创建人名
	CreateName *string `json:"create_name,omitempty"`

	// 创建人名
	CreateAlias *string `json:"create_alias,omitempty"`

	// 任务类型
	TaskType *string `json:"task_type,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 总耗时
	Duration *int64 `json:"duration,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`
}

func (o AlarmHandleHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHandleHistory struct{}"
	}

	return strings.Join([]string{"AlarmHandleHistory", string(data)}, " ")
}
