package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeSeriesData 时间序列数据
type TimeSeriesData struct {

	// 13位时间戳
	Time *int64 `json:"time,omitempty"`

	// 数据值，单位：Kbps（查询流量时）、次（查询事件时）
	Value *float64 `json:"value,omitempty"`
}

func (o TimeSeriesData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeSeriesData struct{}"
	}

	return strings.Join([]string{"TimeSeriesData", string(data)}, " ")
}
