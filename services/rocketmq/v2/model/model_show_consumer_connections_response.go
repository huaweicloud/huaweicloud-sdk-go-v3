package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsumerConnectionsResponse Response Object
type ShowConsumerConnectionsResponse struct {

	// 消费组名称。
	GroupName *string `json:"group_name,omitempty"`

	// 消费组是否在线。
	Online *bool `json:"online,omitempty"`

	// 订阅关系是否一致。
	SubscriptionConsistency *bool `json:"subscription_consistency,omitempty"`

	// 消费者总数。
	Total *int32 `json:"total,omitempty"`

	// 下个分页的offset。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// 上个分页的offset。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// 消费者订阅详情列表。
	Clients        *[]ClientData `json:"clients,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowConsumerConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ShowConsumerConnectionsResponse", string(data)}, " ")
}
