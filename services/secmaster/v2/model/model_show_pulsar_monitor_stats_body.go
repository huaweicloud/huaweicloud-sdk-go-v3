package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowPulsarMonitorStatsBody struct {

	// 毫秒时间戳
	StartTimestamp *int64 `json:"start_timestamp,omitempty"`

	// 毫秒时间戳
	EndTimestamp *int64 `json:"end_timestamp,omitempty"`
}

func (o ShowPulsarMonitorStatsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPulsarMonitorStatsBody struct{}"
	}

	return strings.Join([]string{"ShowPulsarMonitorStatsBody", string(data)}, " ")
}
