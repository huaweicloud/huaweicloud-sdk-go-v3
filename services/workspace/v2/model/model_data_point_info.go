package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataPointInfo 该条告警记录的资源监控数据上报时间和监控数值
type DataPointInfo struct {

	// 计算出该条告警记录的资源监控数据上报的UTC时间
	Time *string `json:"time,omitempty"`

	// 计算出该条告警记录的资源监控数据在该时间点的监控数值
	Value *float64 `json:"value,omitempty"`
}

func (o DataPointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPointInfo struct{}"
	}

	return strings.Join([]string{"DataPointInfo", string(data)}, " ")
}
