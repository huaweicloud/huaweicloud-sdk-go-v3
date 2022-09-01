package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimPricePlanVo struct {

	// 套餐实例id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 账户id
	AccountId *string `json:"account_id,omitempty" xml:"account_id"`

	// sim卡id
	SimCardId *int64 `json:"sim_card_id,omitempty" xml:"sim_card_id"`

	// 套餐状态:0 已删除 1 可激活 2 在用 3. 使用完 4. 可激活 5 已停用 6. 启用失败
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 套餐id
	PricePlanId *string `json:"price_plan_id,omitempty" xml:"price_plan_id"`

	// 容器ID:不同类型卡含义如下 iccid(实体卡)，eid（eSIM）cid（vSIM)
	Cid *string `json:"cid,omitempty" xml:"cid"`

	// 订单id
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 创建时间 例如2020-08-24T07:57:56.000Z
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// 激活时间 例如2020-10-31T16:00:00.000Z
	ActiveTime *sdktime.SdkTime `json:"active_time,omitempty" xml:"active_time"`

	// 停用时间 2021-10-31T16:00:00.000Z
	StopTime *sdktime.SdkTime `json:"stop_time,omitempty" xml:"stop_time"`

	// 总流量(MB)
	FlowTotal *float64 `json:"flow_total,omitempty" xml:"flow_total"`

	// 已使用流量(MB)
	FlowUsed *float64 `json:"flow_used,omitempty" xml:"flow_used"`

	// 剩余流量(MB)
	FlowLeft *float64 `json:"flow_left,omitempty" xml:"flow_left"`

	// 是否使用中(0：否 1：是)
	Using *string `json:"using,omitempty" xml:"using"`

	// 套餐名
	PricePlanName *string `json:"price_plan_name,omitempty" xml:"price_plan_name"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 套餐类型: 0.非自动续费 1.自动续费
	PackageType *int32 `json:"package_type,omitempty" xml:"package_type"`

	// 生效类型: 1.订购后激活使用时生效 2.订购即时生效 3.订购下个月开始生效
	EffectType *int32 `json:"effect_type,omitempty" xml:"effect_type"`

	// 沉默期
	SilentPeriodDay *int32 `json:"silent_period_day,omitempty" xml:"silent_period_day"`

	// 沉默期单位: 1.年 2.月 3.日
	SilentPeriodUnit *int32 `json:"silent_period_unit,omitempty" xml:"silent_period_unit"`

	// 自动续订: 0.不自动续订 1.继续续订
	AutoRenew *int32 `json:"auto_renew,omitempty" xml:"auto_renew"`

	// 位置信息:1.  中国 2.  欧洲 3.  大洋洲 4.  非洲5.  亚太
	LocationDesc *string `json:"location_desc,omitempty" xml:"location_desc"`
}

func (o SimPricePlanVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimPricePlanVo struct{}"
	}

	return strings.Join([]string{"SimPricePlanVo", string(data)}, " ")
}
