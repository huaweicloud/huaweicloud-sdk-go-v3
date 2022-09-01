package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询结果详细。
type SampleDataValue struct {
	Sample *QuerySample `json:"sample,omitempty" xml:"sample"`

	// 时序数据。
	DataPoints *[]MetricDataPoints `json:"data_points,omitempty" xml:"data_points"`
}

func (o SampleDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SampleDataValue struct{}"
	}

	return strings.Join([]string{"SampleDataValue", string(data)}, " ")
}
