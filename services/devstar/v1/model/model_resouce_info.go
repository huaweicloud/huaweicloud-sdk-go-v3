package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResouceInfo struct {

	// 云服务名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 首页链接。
	HomeLink *string `json:"home_link,omitempty"`

	// 开通链接。
	SubscribeLink *string `json:"subscribe_link,omitempty"`

	// 开通指导。
	SubscribeGuide *string `json:"subscribe_guide,omitempty"`

	// 服务类型。
	Type *string `json:"type,omitempty"`

	// 参考价格。
	ReferencePrice *string `json:"reference_price,omitempty"`

	// 价格详情链接。
	PriceDetailsLink *string `json:"price_details_link,omitempty"`

	// 规格,例如: {\"cpu\" : \"0.5\",\"ram\" : 1GB}。
	Specifications *interface{} `json:"specifications,omitempty"`
}

func (o ResouceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResouceInfo struct{}"
	}

	return strings.Join([]string{"ResouceInfo", string(data)}, " ")
}
