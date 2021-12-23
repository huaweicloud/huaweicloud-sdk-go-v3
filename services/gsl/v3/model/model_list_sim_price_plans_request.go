package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSimPricePlansRequest struct {
	// SIM卡标识，可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取

	SimCardId int64 `json:"sim_card_id"`
	// 是否查实时流量

	RealTime *bool `json:"real_time,omitempty"`
	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数

	Limit *int64 `json:"limit,omitempty"`
	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数

	Offset *int64 `json:"offset,omitempty"`
}

func (o ListSimPricePlansRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimPricePlansRequest struct{}"
	}

	return strings.Join([]string{"ListSimPricePlansRequest", string(data)}, " ")
}
