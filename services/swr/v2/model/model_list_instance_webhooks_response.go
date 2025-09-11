package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceWebhooksResponse Response Object
type ListInstanceWebhooksResponse struct {

	// 触发器列表
	Policies *[]WebhookPolicyDetail `json:"policies,omitempty"`

	// 触发器总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceWebhooksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceWebhooksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceWebhooksResponse", string(data)}, " ")
}
