package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataPointInfo struct {

	// 计算出该条告警记录的资源监控数据上报的UTC时间
	Time *string `json:"time,omitempty"`

	// 计算出该条告警记录的资源监控数据在该时间点的监控数值，如：7.019。
	Value *float64 `json:"value,omitempty"`
}

func (o DataPointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPointInfo struct{}"
	}

	return strings.Join([]string{"DataPointInfo", string(data)}, " ")
}
