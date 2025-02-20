package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimelineInfo 时间轴信息。
type TimelineInfo struct {

	// 时间轴名称。
	Name *string `json:"name,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 操作时间。
	OperationTime *string `json:"operation_time,omitempty"`

	// 用户名称。
	UserName *string `json:"user_name,omitempty"`
}

func (o TimelineInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimelineInfo struct{}"
	}

	return strings.Join([]string{"TimelineInfo", string(data)}, " ")
}
