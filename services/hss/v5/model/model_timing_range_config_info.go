package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimingRangeConfigInfo struct {

	// 时间范围
	TimeRange *string `json:"time_range,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o TimingRangeConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimingRangeConfigInfo struct{}"
	}

	return strings.Join([]string{"TimingRangeConfigInfo", string(data)}, " ")
}
