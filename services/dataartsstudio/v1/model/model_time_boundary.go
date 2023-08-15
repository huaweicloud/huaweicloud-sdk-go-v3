package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeBoundary 时间信息
type TimeBoundary struct {

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o TimeBoundary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeBoundary struct{}"
	}

	return strings.Join([]string{"TimeBoundary", string(data)}, " ")
}
