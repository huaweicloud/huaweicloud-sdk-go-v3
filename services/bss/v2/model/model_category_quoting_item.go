package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

// CategoryQuotingItem 分类报价项折扣信息
type CategoryQuotingItem struct {

	// 报价项ID
	ItemId *string `json:"item_id,omitempty"`

	// 云服务编码
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 云服务名称
	CloudServiceTypeName *string `json:"cloud_service_type_name,omitempty"`

	// 商务资源类型
	CommercialResourceType *string `json:"commercial_resource_type,omitempty"`

	// 资源类型编码
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 资源类型名称
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// SKU族编码
	SkuFamilyCode *string `json:"sku_family_code,omitempty"`

	// SKU族名称
	SkuFamilyName *string `json:"sku_family_name,omitempty"`

	// 归属站点编码
	SiteCode *string `json:"site_code,omitempty"`

	// 区域编码
	RegionCode *string `json:"region_code,omitempty"`

	// 区域名称
	RegionName *string `json:"region_name,omitempty"`

	// 可用区AZ编码
	AzCode *string `json:"az_code,omitempty"`

	// 可用区AZ名称
	AzName *string `json:"az_name,omitempty"`

	// 阶梯编号
	StepNo *string `json:"step_no,omitempty"`

	// 计费模式，ONDEMAND：按需、ONETIME：一次性、DAILY：包天、MONTHLY：包月、1_YEARLY：包1年、2_YEARLY：包2年、3_YEARLY：包3年、4_YEARLY：包4年、5_YEARLY：包5年、1_YEARLY_RI：包1年预留实例、3_YEARLY_RI：包3年预留实例
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 折扣率
	DiscountRatio *decimal.Decimal `json:"discount_ratio,omitempty"`

	// 报价项生效时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 报价项失效时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o CategoryQuotingItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryQuotingItem struct{}"
	}

	return strings.Join([]string{"CategoryQuotingItem", string(data)}, " ")
}
