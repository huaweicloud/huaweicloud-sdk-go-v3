package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackPoolsRequest Request Object
type ListBackPoolsRequest struct {

	// 流量池名称
	PoolName *string `json:"pool_name,omitempty"`

	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数
	Offset *int64 `json:"offset,omitempty"`

	// 账期，例如：2021-04
	BillingCycle *string `json:"billing_cycle,omitempty"`

	// 是否查询近六个月账期标识
	AllBillingCycle *bool `json:"all_billing_cycle,omitempty"`
}

func (o ListBackPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListBackPoolsRequest", string(data)}, " ")
}
