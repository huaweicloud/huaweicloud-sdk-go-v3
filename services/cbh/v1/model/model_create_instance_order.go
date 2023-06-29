package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceOrder 生成订单接口请求参数。
type CreateInstanceOrder struct {

	// 云堡垒机实例key。
	InstanceKey int32 `json:"instance_key"`

	// 云服务类型，默认填写“hws.service.type.cbh”。
	CloudServiceType string `json:"cloud_service_type"`

	// 云堡垒机实例局点ID。 可参考[地区和终端节点](https://developer.huaweicloud.com/endpoint)获取
	RegionId string `json:"region_id"`

	// 计费模式。 - 0 包周期计费。
	ChargingMode int32 `json:"charging_mode"`

	// 订购周期类型。 - 2：月 - 3：年
	PeriodType int32 `json:"period_type"`

	// 订购周期数。 - period_type=2（周期类型为月），取值范围[1，9] - periodType=3（周期类型为年），取值范围[1，10]
	PeriodNum int32 `json:"period_num"`

	// 产品信息
	ProductInfos []ProductInfos `json:"product_infos"`

	// 是否自动续订。 - 1，自动续订 - 0，不自动续订（默认）
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 订购数量。当前仅支持订购1个，填写默认值1。
	SubscriptionNum int32 `json:"subscription_num"`
}

func (o CreateInstanceOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceOrder struct{}"
	}

	return strings.Join([]string{"CreateInstanceOrder", string(data)}, " ")
}
