package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolMemVo struct {
	// 流量池标识

	Id *int64 `json:"id,omitempty"`
	// 容器ID

	Cid *string `json:"cid,omitempty"`
	// 套餐订购实例ID

	SimPricePlanId *int64 `json:"sim_price_plan_id,omitempty"`
	// 已用流量(查询账期所在月份), 单位MB

	FlowUsed *float64 `json:"flow_used,omitempty"`
}

func (o PoolMemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMemVo struct{}"
	}

	return strings.Join([]string{"PoolMemVo", string(data)}, " ")
}
