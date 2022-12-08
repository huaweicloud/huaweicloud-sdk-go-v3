package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 生成订单接口请求参数
type CreateInstanceOrder struct {

	// 实例的instance_id
	InstanceKey int32 `json:"instance_key"`

	// CBH云服务类型
	CloudServiceType string `json:"cloud_service_type"`

	// region id
	RegionId string `json:"region_id"`

	// 计费模式
	ChargingMode int32 `json:"charging_mode"`

	// 订购周期类型
	PeriodType int32 `json:"period_type"`

	// 订购周期数
	PeriodNum int32 `json:"period_num"`

	// 订购截止日期
	EndTime *string `json:"end_time,omitempty"`

	// 产品信息
	ProductInfos []ProductInfos `json:"product_infos"`

	// 是否支持续费0/1
	IsAutoRenew int32 `json:"is_auto_renew"`

	// 订购数量
	SubscriptionNum int32 `json:"subscription_num"`

	// 扩容新添
	RelativeResourceId *string `json:"relative_resource_id,omitempty"`
}

func (o CreateInstanceOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceOrder struct{}"
	}

	return strings.Join([]string{"CreateInstanceOrder", string(data)}, " ")
}
