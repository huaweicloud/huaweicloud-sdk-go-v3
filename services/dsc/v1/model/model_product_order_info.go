package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductOrderInfo struct {

	// 租户ID
	TenantId *string `json:"tenantId,omitempty"`

	// 订购周期类型
	PeriodType *string `json:"periodType,omitempty"`

	// 订购周期数量
	PeriodNum *int32 `json:"periodNum,omitempty"`

	// 资源ID
	ResourceId *string `json:"resourceId,omitempty"`

	ProductInfo *ProductInfo `json:"productInfo,omitempty"`
}

func (o ProductOrderInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductOrderInfo struct{}"
	}

	return strings.Join([]string{"ProductOrderInfo", string(data)}, " ")
}
