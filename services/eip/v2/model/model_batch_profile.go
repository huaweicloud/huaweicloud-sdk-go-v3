package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchProfile 订单和产品信息
type BatchProfile struct {

	// 租户id
	UserId *string `json:"user_id,omitempty"`

	// 产品id
	ProductId *string `json:"product_id,omitempty"`

	// 局点id
	RegionId *string `json:"region_id,omitempty"`

	// 订单id
	OrderId *string `json:"order_id,omitempty"`
}

func (o BatchProfile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchProfile struct{}"
	}

	return strings.Join([]string{"BatchProfile", string(data)}, " ")
}
