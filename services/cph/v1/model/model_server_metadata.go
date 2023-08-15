package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerMetadata 订单以及产品相关信息
type ServerMetadata struct {

	// 订单ID，不超过64个字节
	ProductId *string `json:"product_id,omitempty"`

	// 产品ID，不超过64个字节
	OrderId *string `json:"order_id,omitempty"`
}

func (o ServerMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerMetadata struct{}"
	}

	return strings.Join([]string{"ServerMetadata", string(data)}, " ")
}
