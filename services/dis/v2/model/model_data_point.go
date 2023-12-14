package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataPoint 监控数据。
type DataPoint struct {

	// 时间戳。
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 时间戳对应的监控值。
	Value *int64 `json:"value,omitempty"`
}

func (o DataPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPoint struct{}"
	}

	return strings.Join([]string{"DataPoint", string(data)}, " ")
}
