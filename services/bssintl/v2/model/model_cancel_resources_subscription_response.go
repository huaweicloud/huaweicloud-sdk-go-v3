package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CancelResourcesSubscriptionResponse struct {
	// 客户退订订单ID的列表信息。

	OrderIds       *[]string `json:"order_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CancelResourcesSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelResourcesSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"CancelResourcesSubscriptionResponse", string(data)}, " ")
}
