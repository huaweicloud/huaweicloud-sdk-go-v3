package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimPricePlanInfoVo struct {
	// 套餐实例id

	Id *int64 `json:"id,omitempty"`
	// 账户id

	AccountId *string `json:"account_id,omitempty"`
	// sim卡id

	SimCardId *int64 `json:"sim_card_id,omitempty"`
	// 套餐状态:0 已删除 1 可激活 2 在用 3. 使用完 4. 可激活 5 已停用 6. 启用失败

	Status *int32 `json:"status,omitempty"`
	// 套餐id

	PricePlanId *string `json:"price_plan_id,omitempty"`
	// 容器ID:不同类型卡含义如下 iccid(实体卡)，eid（eSIM）cid（vSIM)

	Cid *string `json:"cid,omitempty"`
	// 订单id

	OrderId *string `json:"order_id,omitempty"`
	// 创建时间 例如2020-08-24T07:57:56.000Z

	CreateTime *string `json:"create_time,omitempty"`
	// 激活时间 例如2020-10-31T16:00:00.000Z

	ActiveTime *string `json:"active_time,omitempty"`
	// 停用时间 2021-10-31T16:00:00.000Z

	StopTime *string `json:"stop_time,omitempty"`
	// 总流量(MB)

	FlowTotal *float64 `json:"flow_total,omitempty"`
	// 已使用流量(MB)

	FlowUsed *float64 `json:"flow_used,omitempty"`
	// 剩余流量(MB)

	FlowLeft *float64 `json:"flow_left,omitempty"`
	// 是否使用中(0：否 1：是)

	Using *string `json:"using,omitempty"`
	// 套餐名

	PricePlanName *string `json:"price_plan_name,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 套餐类型: 0.非自动续费 1.自动续费

	PackageType *int32 `json:"package_type,omitempty"`
	// 生效类型: 1.订购后激活使用时生效 2.订购即时生效 3.订购下个月开始生效

	EffectType *int32 `json:"effect_type,omitempty"`
	// 沉默期

	SilentPeriodDay *int32 `json:"silent_period_day,omitempty"`
	// 沉默期单位: 1.年 2.月 3.日

	SilentPeriodUnit *int32 `json:"silent_period_unit,omitempty"`
	// 自动续订: 0.不自动续订 1.继续续订

	AutoRenew *int32 `json:"auto_renew,omitempty"`
	// 位置信息:1.  中国 2.  欧洲 3.  大洋洲 4.  非洲5.  亚太

	LocationDesc *string `json:"location_desc,omitempty"`
}

func (o SimPricePlanInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimPricePlanInfoVo struct{}"
	}

	return strings.Join([]string{"SimPricePlanInfoVo", string(data)}, " ")
}
