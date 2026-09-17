package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdsSlowLogTrendResponse Response Object
type ShowDdsSlowLogTrendResponse struct {

	// 慢日志趋势数量列表
	Points *[]SlowLogPoint `json:"points,omitempty"`

	// 时间间隔
	Interval       *int64 `json:"interval,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDdsSlowLogTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdsSlowLogTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowDdsSlowLogTrendResponse", string(data)}, " ")
}
