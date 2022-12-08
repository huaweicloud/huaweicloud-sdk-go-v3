package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodOrderRequest struct {

	// 计费模式，0：包周期计费，1：按需计费，2：一次性计费
	ChargingMode *int32 `json:"chargingMode,omitempty"`

	// 云服务类型
	CloudServiceType *string `json:"cloudServiceType,omitempty"`

	// 组合套餐ID
	CompositeProductId *string `json:"compositeProductId,omitempty"`

	// 折扣ID
	DiscountId *string `json:"discountId,omitempty"`

	// 是否自动续费
	IsAutoRenew *int32 `json:"isAutoRenew,omitempty"`

	// 订购周期数目
	PeriodNum *int32 `json:"periodNum,omitempty"`

	// 订购周期类型，2：月，3：年
	PeriodType *int32 `json:"periodType,omitempty"`

	// 产品信息列表
	ProductInfos *[]ProductInfoBean `json:"productInfos,omitempty"`

	// 促销ID
	PromotionActivityId *string `json:"promotionActivityId,omitempty"`

	// 促销信息
	PromotionInfo *string `json:"promotionInfo,omitempty"`

	// 当前项目所在region的id，如：xx-xx-1。
	RegionId *string `json:"regionId,omitempty"`

	// 所属国家区域
	Zone *string `json:"zone,omitempty"`
}

func (o PeriodOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodOrderRequest struct{}"
	}

	return strings.Join([]string{"PeriodOrderRequest", string(data)}, " ")
}
