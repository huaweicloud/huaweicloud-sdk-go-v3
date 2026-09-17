package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

// ProductQuotingItem 产品报价项折扣信息
type ProductQuotingItem struct {

	// 报价项ID
	ItemId *string `json:"item_id,omitempty"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty"`

	// 产品规格名称
	ProductSpecName *string `json:"product_spec_name,omitempty"`

	// 云服务编码
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 云服务名称
	CloudServiceTypeName *string `json:"cloud_service_type_name,omitempty"`

	// 运营站点编码
	SiteCode *string `json:"site_code,omitempty"`

	// 产品关联的云服务区信息列表
	RelatedRegions *[]RegionInfo `json:"related_regions,omitempty"`

	// 计费事件编码
	ChargeEventCode *string `json:"charge_event_code,omitempty"`

	// 计费模式，ONDEMAND：按需、ONETIME：一次性、DAILY：包天、MONTHLY：包月、1_YEARLY：包1年、2_YEARLY：包2年、3_YEARLY：包3年、4_YEARLY：包4年、5_YEARLY：包5年、1_YEARLY_RI：包1年预留实例、3_YEARLY_RI：包3年预留实例
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 优惠分类：0：普通优惠，1：产品阶梯，2：分时优惠
	PreferentialCategory *int32 `json:"preferential_category,omitempty"`

	// 优惠方式：0：产品折扣，1：固定单价
	PreferentialType *int32 `json:"preferential_type,omitempty"`

	// 固定单价（preferential_type=1固定单价时有值）
	SalesPrice *decimal.Decimal `json:"sales_price,omitempty"`

	// 折扣率（preferential_type=0产品折扣时有值）
	DiscountRatio *decimal.Decimal `json:"discount_ratio,omitempty"`

	// 计费单位
	PricingBasis *string `json:"pricing_basis,omitempty"`

	// 报价项生效时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 报价项失效时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	ExpireTime *string `json:"expire_time,omitempty"`

	// 产品报价项阶梯列表，产品阶梯或分时优惠时有值返回，否则返回空列表
	ProductQuotingItemSteps *[]ProductQuotingItemStep `json:"product_quoting_item_steps,omitempty"`
}

func (o ProductQuotingItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductQuotingItem struct{}"
	}

	return strings.Join([]string{"ProductQuotingItem", string(data)}, " ")
}
