package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscriptionTasksResponse Response Object
type ListSubscriptionTasksResponse struct {

	// 订阅任务总数
	Total *int32 `json:"total,omitempty"`

	Data           *[]SubscriptionTask `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSubscriptionTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscriptionTasksResponse struct{}"
	}

	return strings.Join([]string{"ListSubscriptionTasksResponse", string(data)}, " ")
}
