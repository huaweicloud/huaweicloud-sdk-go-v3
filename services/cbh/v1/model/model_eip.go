package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性IP
type Eip struct {

	// eip的类型
	IpType string `json:"ip_type"`

	// IP地址对应产品ID
	IpProductId string `json:"ip_product_id"`

	Bandwidth *BandWidth `json:"bandwidth"`
}

func (o Eip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Eip struct{}"
	}

	return strings.Join([]string{"Eip", string(data)}, " ")
}
