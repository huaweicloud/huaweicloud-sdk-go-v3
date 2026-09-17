package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogTrendNewResponse Response Object
type ShowSlowLogTrendNewResponse struct {

	// 趋势数量列表
	TrendData *[]SlowLogTrendPoint `json:"trend_data,omitempty"`

	// 时间间隔
	Interval       *int64 `json:"interval,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSlowLogTrendNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogTrendNewResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogTrendNewResponse", string(data)}, " ")
}
