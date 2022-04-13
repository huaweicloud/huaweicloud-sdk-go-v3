package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Profile对象
type ProfileResp struct {
	// 订单的id

	OrderId *string `json:"order_id,omitempty"`
	// 产品的id

	ProductId *string `json:"product_id,omitempty"`
	// region的id

	RegionId *string `json:"region_id,omitempty"`
	// 用户的id

	UserId *string `json:"user_id,omitempty"`
}

func (o ProfileResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProfileResp struct{}"
	}

	return strings.Join([]string{"ProfileResp", string(data)}, " ")
}
