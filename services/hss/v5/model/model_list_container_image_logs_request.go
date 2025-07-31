package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerImageLogsRequest Request Object
type ListContainerImageLogsRequest struct {

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	// 事件名称
	EventName *string `json:"event_name,omitempty"`

	// 触发事件的源ip
	SourceIp *string `json:"source_ip,omitempty"`

	// 操作用户
	UserName *string `json:"user_name,omitempty"`

	// 查询日志范围的最小时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询日志范围的最大时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 每页显示个数，默认为10
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListContainerImageLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerImageLogsRequest struct{}"
	}

	return strings.Join([]string{"ListContainerImageLogsRequest", string(data)}, " ")
}
