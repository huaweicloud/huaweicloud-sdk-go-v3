package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceEmergencyLogsRequest Request Object
type ListInstanceEmergencyLogsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 开始时间（Unix时间戳，毫秒）
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间（Unix时间戳，毫秒）
	EndTime *int64 `json:"end_time,omitempty"`

	// 当前页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页记录数
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListInstanceEmergencyLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceEmergencyLogsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceEmergencyLogsRequest", string(data)}, " ")
}
