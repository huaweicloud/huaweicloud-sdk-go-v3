package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OrderLineItemEntityV2 struct {

	// 订单项ID。
	OrderLineItemId *string `json:"order_line_item_id,omitempty"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 云服务类型名称。例如ECS的云服务类型名称为“弹性云服务器”。
	ServiceTypeName *string `json:"service_type_name,omitempty"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 产品规格描述。
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`

	// 周期类型。 0：天1：周2：月3：年4：小时5：一次性6：按需（预留）7：按用量报表使用（预留）
	PeriodType *int32 `json:"period_type,omitempty"`

	// 周期数量。  说明： 当订单为退订资源的订单时，参数取值为null。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 订购数量。
	SubscriptionNum *int32 `json:"subscription_num,omitempty"`

	// 订单优惠后金额（实付价格，不含券不含卡）。
	AmountAfterDiscount *float64 `json:"amount_after_discount,omitempty"`

	// 订单金额（官网价）。 退订订单中，该金额等于currencyAfterDiscount。
	OfficialAmount *float64 `json:"official_amount,omitempty"`

	AmountInfo *AmountInfomationV2 `json:"amount_info,omitempty"`

	// 货币编码。
	Currency *string `json:"currency,omitempty"`

	// 产品目录编码。
	CategoryCode *string `json:"category_code,omitempty"`

	// 产品归属的云服务类型编码。 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。
	ProductOwnerService *string `json:"product_owner_service,omitempty"`

	// 商务归属的资源类型编码。 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。
	CommercialResource *string `json:"commercial_resource,omitempty"`

	BaseProductInfo *ProductObject `json:"base_product_info,omitempty"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`
}

func (o OrderLineItemEntityV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderLineItemEntityV2 struct{}"
	}

	return strings.Join([]string{"OrderLineItemEntityV2", string(data)}, " ")
}
