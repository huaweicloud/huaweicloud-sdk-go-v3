package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudPhoneServerDetailResponseBodyMetadata 订单以及产品相关信息。
type ShowCloudPhoneServerDetailResponseBodyMetadata struct {

	// 产品ID，不超过64个字节。
	ProductId *string `json:"product_id,omitempty"`

	// 订单ID，不超过64个字节。
	OrderId *string `json:"order_id,omitempty"`
}

func (o ShowCloudPhoneServerDetailResponseBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneServerDetailResponseBodyMetadata struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneServerDetailResponseBodyMetadata", string(data)}, " ")
}
