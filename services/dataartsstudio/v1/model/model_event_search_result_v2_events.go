package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventSearchResultV2Events struct {

	// 事件ID。
	Id *string `json:"id,omitempty"`

	// 事件名称。
	Name *string `json:"name,omitempty"`

	// 事件类型，取值为ERROR和SLOW_DOWN。
	Type *string `json:"type,omitempty"`

	// 事件状态，取值为NEW_DISCOVERY、PROCESSING、RESTORED和IGNORED。
	Status *string `json:"status,omitempty"`

	// 基线任务名称。
	BaselineName *string `json:"baseline_name,omitempty"`

	// 作业名称。
	TaskName *string `json:"task_name,omitempty"`

	// 作业ID。
	TaskId *string `json:"task_id,omitempty"`

	// 作业版本号。
	TaskVersion *int32 `json:"task_version,omitempty"`

	// 发生时间戳，单位毫秒。
	HappenTime *int32 `json:"happen_time,omitempty"`

	// 责任人用户名称。
	OwnerName *string `json:"owner_name,omitempty"`
}

func (o EventSearchResultV2Events) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventSearchResultV2Events struct{}"
	}

	return strings.Join([]string{"EventSearchResultV2Events", string(data)}, " ")
}
