package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeRangeItem period每日生效时间区间
type TimeRangeItem struct {

	// period每日生效时间区间开始
	St *int32 `json:"st,omitempty"`

	// period每日生效时间区间结束
	End *int32 `json:"end,omitempty"`
}

func (o TimeRangeItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeRangeItem struct{}"
	}

	return strings.Join([]string{"TimeRangeItem", string(data)}, " ")
}
