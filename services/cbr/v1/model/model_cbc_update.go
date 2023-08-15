package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CbcUpdate struct {

	// 云服务ConsoleURL。订单支付完成后，客户可以通过此URL跳转到云服务Console页面查看信息
	CloudServiceConsoleURL *string `json:"cloudServiceConsoleURL,omitempty"`

	ProductInfo *CbcProductInfoUpdate `json:"productInfo"`

	// 待变更的资源ID
	ResourceId string `json:"resourceId"`

	// 购买折扣
	PromotionInfo *string `json:"promotion_info,omitempty"`
}

func (o CbcUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcUpdate struct{}"
	}

	return strings.Join([]string{"CbcUpdate", string(data)}, " ")
}
