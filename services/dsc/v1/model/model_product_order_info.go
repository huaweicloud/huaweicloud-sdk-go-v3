package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductOrderInfo struct {

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 订购周期类型
	PeriodType *string `json:"period_type,omitempty"`

	// 订购周期数量
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 产品信息
	ProductInfo *[]ProductInfoBean `json:"product_info,omitempty"`
}

func (o ProductOrderInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductOrderInfo struct{}"
	}

	return strings.Join([]string{"ProductOrderInfo", string(data)}, " ")
}
