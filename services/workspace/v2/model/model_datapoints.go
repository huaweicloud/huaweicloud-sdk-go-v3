package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Datapoints 桌面监控Cpu与Mem时间点信息。
type Datapoints struct {

	// 平均数值。
	Average *float64 `json:"average,omitempty"`

	// 时间戳。
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o Datapoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Datapoints struct{}"
	}

	return strings.Join([]string{"Datapoints", string(data)}, " ")
}
