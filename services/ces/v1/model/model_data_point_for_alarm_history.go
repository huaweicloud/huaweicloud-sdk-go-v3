package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 计算出该条告警历史的资源监控数据上报时间和监控数值。
type DataPointForAlarmHistory struct {

	// 计算出该条告警历史的资源监控数据上报时间，UNIX时间戳，单位毫秒，如：1603131028000。
	Time *int64 `json:"time,omitempty" xml:"time"`

	// 计算出该条告警历史的资源监控数据在该时间点的监控数值，如：7.019。
	Value *float64 `json:"value,omitempty" xml:"value"`
}

func (o DataPointForAlarmHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPointForAlarmHistory struct{}"
	}

	return strings.Join([]string{"DataPointForAlarmHistory", string(data)}, " ")
}
