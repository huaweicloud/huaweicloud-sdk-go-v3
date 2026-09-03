package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockTrendResponse Response Object
type ShowDeadLockTrendResponse struct {

	// 时间间隔（ms）
	IntervalMillis *int64 `json:"interval_millis,omitempty"`

	// 趋势数量列表
	TrendList      *[]DeadLockTrendPoint `json:"trend_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowDeadLockTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTrendResponse", string(data)}, " ")
}
