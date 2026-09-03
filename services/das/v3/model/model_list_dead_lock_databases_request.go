package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeadLockDatabasesRequest Request Object
type ListDeadLockDatabasesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 开始时间戳 ms
	StartTime int64 `json:"start_time"`

	// 结束时间戳 ms
	EndTime int64 `json:"end_time"`
}

func (o ListDeadLockDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeadLockDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListDeadLockDatabasesRequest", string(data)}, " ")
}
