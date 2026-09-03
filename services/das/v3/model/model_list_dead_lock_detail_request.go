package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeadLockDetailRequest Request Object
type ListDeadLockDetailRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库名称
	DbName string `json:"db_name"`

	// 开始时间戳 ms
	StartTime int64 `json:"start_time"`

	// 结束时间戳 ms
	EndTime int64 `json:"end_time"`

	// 页码
	CurPage int32 `json:"cur_page"`

	// 每页记录数
	PerPage int32 `json:"per_page"`
}

func (o ListDeadLockDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeadLockDetailRequest struct{}"
	}

	return strings.Join([]string{"ListDeadLockDetailRequest", string(data)}, " ")
}
