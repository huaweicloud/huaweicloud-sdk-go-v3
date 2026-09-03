package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockStatisticsRequest Request Object
type ShowDeadLockStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 当前时间戳 ms
	CurrentTime int64 `json:"current_time"`

	// 开始时间戳 ms
	StartTime int64 `json:"start_time"`

	// 结束时间戳 ms
	EndTime int64 `json:"end_time"`
}

func (o ShowDeadLockStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowDeadLockStatisticsRequest", string(data)}, " ")
}
