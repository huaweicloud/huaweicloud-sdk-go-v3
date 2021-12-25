package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 产品
type ProductReferer struct {
	// 产品ID，未填写厂商ID+型号时产品ID必填

	ProductId *int32 `json:"product_id,omitempty"`
	// 产品名称

	ProductName *string `json:"product_name,omitempty"`
	// 厂商ID，未填写产品ID时厂商ID和型号必填

	ManufacturerId *string `json:"manufacturer_id,omitempty"`
	// 型号，未填写产品ID时厂商ID和型号必填

	Model *string `json:"model,omitempty"`
	// 产品的协议类型：0-mqtt，1-coap，2-modbus，3-http, 4-opcua

	ProtocolType *int32 `json:"protocol_type,omitempty"`
	// 产品类型：0-普通产品 1-网关产品

	ProductType *int32 `json:"product_type,omitempty"`
}

func (o ProductReferer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductReferer struct{}"
	}

	return strings.Join([]string{"ProductReferer", string(data)}, " ")
}
