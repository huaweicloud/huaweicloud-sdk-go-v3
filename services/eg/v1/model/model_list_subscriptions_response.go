package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionsResponse Response Object
type ListSubscriptionsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]SubscriptionInfo `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSubscriptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionsResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionsResponse", string(data)}, " ")
}
