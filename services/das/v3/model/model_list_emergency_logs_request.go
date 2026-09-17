package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEmergencyLogsRequest Request Object
type ListEmergencyLogsRequest struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 开始时间（Unix timestamp，毫秒）
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间（Unix timestamp，毫秒）
	EndTime *int64 `json:"end_time,omitempty"`

	// 页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListEmergencyLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEmergencyLogsRequest struct{}"
	}

	return strings.Join([]string{"ListEmergencyLogsRequest", string(data)}, " ")
}
