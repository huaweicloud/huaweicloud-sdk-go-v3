package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackPoolMemVo struct {

	// 流量池标识
	Id *int64 `json:"id,omitempty"`

	// 容器ID
	Cid *string `json:"cid,omitempty"`

	// 套餐订购实例ID
	SimPricePlanId *int64 `json:"sim_price_plan_id,omitempty"`

	// 已用流量(查询账期所在月份), 单位MB
	FlowUsed *float64 `json:"flow_used,omitempty"`

	// 卡当前状态：10-可测试，11-未激活，13-可激活，14-已停用，20-在用，30-已拆机
	SimStatus *int32 `json:"sim_status,omitempty"`
}

func (o BackPoolMemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackPoolMemVo struct{}"
	}

	return strings.Join([]string{"BackPoolMemVo", string(data)}, " ")
}
