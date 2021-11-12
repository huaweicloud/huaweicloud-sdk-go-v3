package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSimPoolMembersRequest struct {
	// 流量池标识

	SimPoolId int64 `json:"sim_pool_id"`
	// 容器ID

	Cid *string `json:"cid,omitempty"`
	// 每页记录数

	Limit *int64 `json:"limit,omitempty"`
	// 页码

	Offset *int64 `json:"offset,omitempty"`
	// 账期，例如：2021-04

	BillingCycle string `json:"billing_cycle"`
}

func (o ListSimPoolMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimPoolMembersRequest struct{}"
	}

	return strings.Join([]string{"ListSimPoolMembersRequest", string(data)}, " ")
}
