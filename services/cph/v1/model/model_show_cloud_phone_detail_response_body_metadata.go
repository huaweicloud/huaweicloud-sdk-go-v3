package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 订单以及产品相关信息
type ShowCloudPhoneDetailResponseBodyMetadata struct {

	// 订单ID，不超过64个字节
	OrderId *string `json:"order_id,omitempty"`

	// 产品ID，不超过64个字节
	ProductId *string `json:"product_id,omitempty"`
}

func (o ShowCloudPhoneDetailResponseBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneDetailResponseBodyMetadata struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneDetailResponseBodyMetadata", string(data)}, " ")
}
