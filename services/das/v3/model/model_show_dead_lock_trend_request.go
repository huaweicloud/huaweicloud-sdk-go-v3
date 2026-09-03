package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockTrendRequest Request Object
type ShowDeadLockTrendRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 开始时间戳 ms
	StartTime int64 `json:"start_time"`

	// 结束时间戳 ms
	EndTime int64 `json:"end_time"`
}

func (o ShowDeadLockTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTrendRequest", string(data)}, " ")
}
