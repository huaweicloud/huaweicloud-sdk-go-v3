package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBusinessDiscountInfoResponse Response Object
type ListBusinessDiscountInfoResponse struct {

	// 总条数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 商务ID
	CommerceId *string `json:"commerce_id,omitempty"`

	// 商务编号
	CommerceCode *string `json:"commerce_code,omitempty"`

	// 商务生效时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 商务失效时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	ExpireTime *string `json:"expire_time,omitempty"`

	// 产品报价项列表（quoting_item_type=PRODUCT_ITEM时有值返回，否则返回空列表）
	ProductQuotingItems *[]ProductQuotingItem `json:"product_quoting_items,omitempty"`

	// 分类报价项列表（quoting_item_type=CATEGORY_ITEM时有值返回，否则返回空列表）
	CategoryQuotingItems *[]CategoryQuotingItem `json:"category_quoting_items,omitempty"`

	// 分类报价项阶梯列表（quoting_item_type=CATEGORY_ITEM时有值返回，否则返回空列表）
	CategoryQuotingItemSteps *[]CategoryQuotingItemStep `json:"category_quoting_item_steps,omitempty"`

	// 阶梯累计周期类型，category_quoting_item_steps有值返回时返回
	AccumulationCycleType *string `json:"accumulation_cycle_type,omitempty"`

	// 运营站点列表
	Sites          *[]SiteInfo `json:"sites,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListBusinessDiscountInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessDiscountInfoResponse struct{}"
	}

	return strings.Join([]string{"ListBusinessDiscountInfoResponse", string(data)}, " ")
}
