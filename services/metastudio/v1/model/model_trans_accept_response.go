package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransAcceptResponse struct {

	// 资产转移时，是否需要从接收方扣减资源（产生计费话单)
	IsNeedBilling *bool `json:"is_need_billing,omitempty"`

	// 需要扣减的资源列表。
	RequiredResources *[]BillResources `json:"required_resources,omitempty"`
}

func (o TransAcceptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransAcceptResponse struct{}"
	}

	return strings.Join([]string{"TransAcceptResponse", string(data)}, " ")
}
