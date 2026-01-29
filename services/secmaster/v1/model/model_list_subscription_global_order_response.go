package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionGlobalOrderResponse Response Object
type ListSubscriptionGlobalOrderResponse struct {

	// 已购资源列表
	Resources      *[]SubscriptionGlobalResource `json:"resources,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListSubscriptionGlobalOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionGlobalOrderResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionGlobalOrderResponse", string(data)}, " ")
}
