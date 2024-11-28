package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CbcOrderChange struct {

	// 云服务ConsoleURL。订单支付完成后，客户可以通过此URL跳转到云服务Console页面查看信息
	CloudServiceConsoleUrl *string `json:"cloud_service_console_url,omitempty"`

	ProductInfo *CbcProductInfoOrderChange `json:"product_info"`

	// 待变更的资源ID
	ResourceId string `json:"resource_id"`

	// 是否自动支付，默认非自动支付：false
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// 购买折扣
	PromotionInfo *string `json:"promotion_info,omitempty"`
}

func (o CbcOrderChange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcOrderChange struct{}"
	}

	return strings.Join([]string{"CbcOrderChange", string(data)}, " ")
}
