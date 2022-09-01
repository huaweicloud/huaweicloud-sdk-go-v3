package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResouceInfo struct {

	// 云服务名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 首页链接。
	HomeLink *string `json:"home_link,omitempty" xml:"home_link"`

	// 开通链接。
	SubscribeLink *string `json:"subscribe_link,omitempty" xml:"subscribe_link"`

	// 服务类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 参考价格。
	ReferencePrice *string `json:"reference_price,omitempty" xml:"reference_price"`

	// 价格详情链接。
	PriceDetailsLink *string `json:"price_details_link,omitempty" xml:"price_details_link"`

	// 规格,例如: {\"cpu\" : \"0.5\",\"ram\" : 1GB}。
	Specifications *interface{} `json:"specifications,omitempty" xml:"specifications"`
}

func (o ResouceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResouceInfo struct{}"
	}

	return strings.Join([]string{"ResouceInfo", string(data)}, " ")
}
