package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataPoints 云应用服务器监控Cpu与Mem时间点信息。
type DataPoints struct {

	// 平均数值。
	Average *float64 `json:"average,omitempty"`

	// 时间戳。
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o DataPoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPoints struct{}"
	}

	return strings.Join([]string{"DataPoints", string(data)}, " ")
}
