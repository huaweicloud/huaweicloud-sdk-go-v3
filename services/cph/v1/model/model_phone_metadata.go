package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneMetadata 订单以及产品相关信息。
type PhoneMetadata struct {

	// 订单ID，不超过64个字节。
	OrderId *string `json:"order_id,omitempty"`

	// 产品ID，不超过64个字节。
	ProductId *string `json:"product_id,omitempty"`
}

func (o PhoneMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneMetadata struct{}"
	}

	return strings.Join([]string{"PhoneMetadata", string(data)}, " ")
}
