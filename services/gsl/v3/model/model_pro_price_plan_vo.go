package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProPricePlanVo struct {

	// 套餐ID
	PricePlanId *string `json:"price_plan_id,omitempty"`

	// 套餐名称
	PricePlanName *string `json:"price_plan_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 流量总量(MB)
	FlowTotal *int64 `json:"flow_total,omitempty"`

	// 套餐类型 1基础套餐;2叠加包套餐;如果是国际漫游不区分基础套餐包和叠加包
	PackageType *int32 `json:"package_type,omitempty"`

	// 套餐周期
	Period *int32 `json:"period,omitempty"`

	// 套餐周期类型 10:日;20:月;30:季;40:半年;50:年
	PeriodType *int32 `json:"period_type,omitempty"`

	// 套餐生效类型 1.订购后激活使用时生效 2.订购即时生效 3.订购下个月开始生效
	EffectType *int32 `json:"effect_type,omitempty"`

	// 沉默期
	SilentPeriodDay *int32 `json:"silent_period_day,omitempty"`

	// 沉默期单位 1.年 2.月 3.日
	SilentPeriodUnit *int32 `json:"silent_period_unit,omitempty"`

	// 是否自动续订
	AutoRenew *bool `json:"auto_renew,omitempty"`

	// 套餐适用区域
	LocationDesc *string `json:"location_desc,omitempty"`

	// 区域 1.中国 2.欧洲 3.大洋洲 4.非洲 5.亚太
	LocationType *int32 `json:"location_type,omitempty"`

	// SIM卡类型 1.vSIM 2.eSIM 3.实体卡
	SimType *int32 `json:"sim_type,omitempty"`

	// 运营商 101/1 中国移动/中国移动（实体卡） 102/2中国电信/中国电信（实体卡） 3中国联通（实体卡） 201.欧洲 501.中国香港 502.中国澳门 503.泰国 504.日本 505.柬埔寨 506.印度尼西亚 507.马来西亚 508.新加坡 509.斯里兰卡 510.中国台湾 511.孟加拉
	CarrierType *int32 `json:"carrier_type,omitempty"`

	// 价格(分)
	Price *int32 `json:"price,omitempty"`
}

func (o ProPricePlanVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProPricePlanVo struct{}"
	}

	return strings.Join([]string{"ProPricePlanVo", string(data)}, " ")
}
