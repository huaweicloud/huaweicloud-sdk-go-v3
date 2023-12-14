package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Metrics 数据对象。
type Metrics struct {

	// 监控数据。
	DataPoints *[]DataPoint `json:"dataPoints,omitempty"`

	// 监控指标。
	Label *string `json:"label,omitempty"`
}

func (o Metrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metrics struct{}"
	}

	return strings.Join([]string{"Metrics", string(data)}, " ")
}
